<?php
namespace App\Entity;

class UserEntity
{

    private int $settingsId;

    private string $key;

    private string $value;

    private int $parent;

    private int $type;

    private string $description;


    public function getSettingsId(): int
    {
        return $this->settingsId;
    }

    public function getKey(): string
    {
        return $this->key;
    }

    public function getValue(): string
    {
        return $this->value;
    }

    public function getParent(): int
    {
        return $this->parent;
    }

    public function getType(): int
    {
        return $this->type;
    }

    public function getDescription(): string
    {
        return $this->description;
    }


    public function toArray(): array
    {
        return [
            'settingsId' => $this->settingsId,
            'key' => $this->key,
            'value' => $this->value,
            'parent' => $this->parent,
            'type' => $this->type,
            'description' => $this->description,
        ];
    }

    public static function fromArray(array $fields): self
    {
        $entity = new self();

        $entity->settingsId = $fields['settingsId'];
        $entity->key = $fields['key'];
        $entity->value = $fields['value'];
        $entity->parent = $fields['parent'];
        $entity->type = $fields['type'];
        $entity->description = $fields['description'];

        return $entity;
    }
}