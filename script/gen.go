package main

import (
	"gorm.io/gen"
	"gorm.io/gen/field"
	"hook007/pkg/db"
	"strings"
)

type (
	Generate struct {
		Gen *gen.Generator
	}

	Relate struct {
		Table       string
		FieldRelate []*FieldRelate
	}

	FieldRelate struct {
		Model        interface{}
		Type         field.RelationshipType
		FieldTagName string
		Table        string
		RelateConfig *field.RelateConfig
	}
)

var (
	relate = []Relate{
		//{
		//	Table: "qw",
		//	FieldRelate: []*FieldRelate{
		//		{
		//			Model:        model_gen.CorpGroup{},
		//			Type:         field.HasMany,
		//			FieldTagName: "CorpGroup",
		//			Table:        "corp_group",
		//			RelateConfig: &field.RelateConfig{
		//				RelateSlicePointer: true, // 指针
		//				GORMTag: field.GormTag{
		//					"foreignKey": []string{"QwID"},
		//				},
		//				JSONTag: "corpGroup",
		//			},
		//		},
		//	},
		//},
	}
)

func ConnectGen(path string) error {
	mysqlDB, err := db.Connect()
	if err != nil {
		return err
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      path + "dao/query",
		ModelPkgPath: path + "dao/model_gen",
		Mode:         gen.WithDefaultQuery | gen.WithoutContext,
	})

	// 使用数据库
	g.UseDB(mysqlDB)

	generate := Generate{
		Gen: g,
	}

	// 生成 dal 关系模型
	generate.Relate(relate)

	// 生成所有数据库
	g.ApplyBasic(g.GenerateAllTable(generate.FieldOpts()...)...)

	// 生成 model_gen 关系模型
	generate.Relate(relate)

	g.Execute()

	return nil
}

func (g *Generate) FieldOpts() []gen.ModelOpt {
	updatedAt := gen.FieldGORMTag("updated_at", func(tag field.GormTag) field.GormTag {
		tag.Set("autoUpdateTime", "milli")
		return tag
	})
	createdAt := gen.FieldGORMTag("created_at", func(tag field.GormTag) field.GormTag {
		tag.Set("autoCreateTime", "milli")
		return tag
	})

	// 自定义模型结体字段的标签
	// 指定 tag 设置 json
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		// 把 json 修改为驼峰式
		if strings.Contains(columnName, "_") {
			parts := strings.Split(columnName, "_")
			for i := range parts {
				if i > 0 {
					parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
				}
			}
			columnName = strings.Join(parts, "")
		}
		// 把更新和删除的 json 输出修改为 -
		neglectCloumn := []string{"deletedAt"}
		for _, v := range neglectCloumn {
			if v == columnName {
				columnName = "-"
				break
			}
		}
		return columnName
	})

	return []gen.ModelOpt{createdAt, updatedAt, jsonField}
}

// Relate 关联模型
func (g *Generate) Relate(relates []Relate) {
	for _, table := range relates {
		opts := g.FieldOpts()
		for _, relate := range table.FieldRelate {
			if relate.Model != nil {
				opts = append(opts, gen.FieldRelateModel(
					relate.Type,
					relate.FieldTagName,
					relate.Model,
					relate.RelateConfig,
				))
			} else {
				genModel := g.Gen.GenerateModel(relate.Table, g.FieldOpts()...)
				opts = append(opts, gen.FieldRelate(
					relate.Type,
					relate.FieldTagName,
					genModel,
					relate.RelateConfig,
				))
			}
		}
		relateModel := g.Gen.GenerateModel(table.Table, opts...)
		g.Gen.ApplyBasic(relateModel)
	}
}
