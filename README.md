# Generator PHP Entities

Run first time
`go run backend/app/main.go --table=settings --db_user=user --db_name=dbname --db_password=pass --db_port=3306 --db_host=127.0.0.1`

Run Second
`main.go --table=users`

Execute:

```php
class UserEntity
{
    private int $id;

    /**
    * @return int
    */
    public function getId(): int
    {
        return $this->id;
    }

    /**
    * @return array
    */
    public function toArray(): array
    {
        return [
            'id' => $this->id;
        ];
    }

    /**
    * @param array $fields
    * @return static
    */
    public static function fromArray(array $fields): self
    {
        $entity = new self();
        $entity->id = $fields['id'];

        return $entity;
    }
}
```

## Generate DTO
Sometimes we need create DTO struct form array

```php
class UserDto
{
    private int $id;

    /**
    * @return int
    */
    public function getId(): int
    {
        return $this->id;
    }

    /**
    * @param array $fields
    * @return static
    */
    public static function fromArray(array $fields): self
    {
        $entity = new self();
        $entity->id = $fields['id'];

        return $entity;
    }
}
```
How convenient is it to add parameters in go for parse?
yaml?
```yaml
entity:
    name: UserEntity
    from: database
    database:
        alias: testbd
        name: db_name
        port: db_port
        user: db_user
        pass: db_password
dto:
    name: UserDto
    from: array
    params:
        id: int,
        name: string
```


## Generate ValuesObject



## Description options
| Options             | Description |
|---------------------|-------------|
| -n, --db_name=      |  DB Name |
| -h, --db_host=      |  DB Host |
| -p, --db_port=      |  DB Port |
| -u, --db_user=      |  DB User |
| --db_password=      | DB Password |
| --db_type=          | Type of DB |
| -y, --type=         |  Type of generates files (default: entity) |
| -t, --table=        |  Table for generate Entity |
| -o, --output_path=  |  Path where generation file(s) are saved |
| -s, --storage_path= | Storage Path (default: /var/tmp/jtrw_generator_php_entities.db) |
| --profile=          | Profile's credentials. Command 'list' for display all profiles |


