<?php
namespace App\Entity;

class UserEntity
{

    private int $id;

    private string $name;

    public function getId(): string
    {
        return $this->id;
    }
}