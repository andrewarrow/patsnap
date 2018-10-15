# patsnap
Dog!

#https://app.datadoghq.com/api/v1/logs-queries/scopes/7156/list-csv?query=

{"list":{"time":{"offset":-7200000,"from":"now-604800s","to":"now"},
         "search":{"query":"@thing:866 AND @otherthing:foo"},
         "columns":[{"field":{"path":"@bar"}},
                    {"field":{"path":"@msg"}}],
         "limit":5000,
         "sort":{"time":{"order":"desc"}}}}
