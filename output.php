<?php
namespace App\Entity;

class UserEntity
{

    private int $id;

    private string $name;


    public function getId(): int
    {
        return $this->id;
    }

    public function getName(): string
    {
        return $this->name;
    }


    public function toArray(): array
    {
        return [
            'id' => $this->id,
            'name' => $this->name,
        ];
    }

    public static function fromArray(array $fields): self
    {
        $entity = new self();

        $entity->id = $fields['id'];
        $entity->name = $fields['name'];

        return $entity;
    }
}