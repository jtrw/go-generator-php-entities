# Generator PHP Entities

Run Frist time
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