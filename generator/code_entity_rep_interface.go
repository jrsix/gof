package generator

var (
	// 实体仓储接口模板
	TPL_ENTITY_REP_INTERFACE = CodeTemplate(
		`// auto generate by gof (http://github.com/jsix/gof)
        package {{.VAR.IRepoPkgName}}

        import(
            "{{.VAR.IRepoPkg}}"
        )

        type I<R> interface{
            // auto generate by gof
            // Get <E>
            Get<R2>_(primary interface{})*<E2>
            // GetBy <E>
            Get<R2>By_(where string,v ...interface{})*<E2>
            // Select <E>
            Select<R2>_(where string,v ...interface{})[]*<E2>
            // Save <E>
            Save<R2>_(v *<E2>)(int,error)
            // Delete <E>
            Delete<R2>_(primary interface{}) error
            // Batch Delete <E>
            BatchDelete<R2>_(where string,v ...interface{})(int64,error)
        }`)
)
