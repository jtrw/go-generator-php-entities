# Generator PHP Entities

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