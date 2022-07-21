# Generator PHP Entities

`main.go --table=users`

Execute:

```php
class UserEntity
{
    private int $id;

    public function getId(): int
    {
        return $this->id;
    }

    ...
}
```