package db

func getMysqlDsn(opts Options) (string)
{
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", opts.DbUser, opts.DbPassword, opts.DbHost, opts.DbPort, opts.DbName)
}

func getPgsqlDsn() (string)
{

}