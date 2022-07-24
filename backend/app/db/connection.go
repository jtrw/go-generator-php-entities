package db

func getMysqlDsn(opts Options) (string)
{
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", opts.DbUser, opts.DbPassword, opts.DbHost, opts.DbPort, opts.DbName)
}

func getPgsqlDsn(opts Options) (string)
{
    return fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        opts.DbHost,
        opts.DbPort,
        opts.DbUser,
        opts.DbPassword,
        opts.DbName)
}