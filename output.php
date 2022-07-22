<?php
namespace App\Entity;

class UserEntity
{

    private int $id;

    private string $name;

    private int $userId;

    private string $hmac;

    private string $uin;

    private string $token;

    private string $tokenDate;

    private string $startedGameToken;

    private int $startedGameId;

    private string $gamesessionid;

    private string $srvSessionId;

    private string $facebookUserId;

    private string $registrationDate;

    private string $registrationPlatform;

    private string $registrationIp;

    private string $registrationCountry;

    private int $domain;

    private string $emailPlatform;

    private string $emailIp;

    private string $lastvisitDate;

    private string $lastvisitIp;

    private int $lastvisitDevice;

    private int $langId;

    private int $loginAttempts;

    private string $password;

    private string $passwordResetCode;

    private string $email;

    private int $emailConfirmed;

    private string $emailConfirmCode;

    private string $emailConfirmDate;

    private string $emailConfirmCountry;

    private int $verificationEmailAttempts;

    private int $isSendNews;

    private int $isSendSms;

    private string $avatar;

    private string $firstName;

    private string $lastName;

    private string $sex;

    private string $birthday;

    private string $phone;

    private int $phoneVerified;

    private string $phoneVerificationCode;

    private string $phoneVerificationDate;

    private int $phoneCallblock;

    private int $verificationPhoneAttempts;

    private int $countryId;

    private string $city;

    private string $address;

    private string $zip;

    private float $rating;

    private int $currencyId;

    private int $affiliateId;

    private string $affiliateIdUpdatedOn;

    private string $subaff;

    private string $subaff1;

    private string $subaff2;

    private string $src;

    private string $media;

    private string $splitPointDate;

    private float $balance;

    private float $balanceBonuses;

    private int $balanceVersion;

    private int $wagerPercent;

    private int $wagerReset;

    private int $noDailyDepositLimit;

    private int $disableDeposits;

    private float $depositsLimitDay;

    private float $depositsLimitWeek;

    private float $depositsLimitMonth;

    private float $minDepositAmount;

    private float $maxDepositAmount;

    private float $minWithdrawAmount;

    private float $maxWithdrawAmount;

    private int $vipId;

    private string $vipDate;

    private int $vipPercent;

    private int $vipPoints;

    private string $totalCalcDate;

    private float $totalBetsCash;

    private float $totalBetsChips;

    private float $totalWinsCash;

    private float $totalWinsChips;

    private float $totalWdsCash;

    private float $totalWdsChips;

    private float $totalDepCash;

    private float $totalBonCash;

    private float $totalBonChips;

    private float $totalBonCancel;

    private float $totalWdsCancel;

    private float $totalWdsApproved;

    private float $totalWdsPending;

    private float $totalBonPay;

    private float $totalDeposits;

    private float $totalWithdrawals;

    private int $documentsProvided;

    private string $documentsProvidedDate;

    private int $documentsProvidedAdmin;

    private int $showInReport;

    private int $frozen;

    private string $modifiedOn;

    private int $blockDeposits;

    private string $zotapayEndpoint;

    private int $paymentTypeId;

    private int $notifications;

    private int $eproRestricted;

    private string $eproRestrictedUpdate;

    private int $whereReg;

    private string $cubitsBitcoinAddress;

    private string $cubitsChannelId;

    private string $coinspaidBitcoinAddress;

    private string $coinspaidBchAddress;

    private string $coinspaidLtcAddress;

    private string $coinspaidEthAddress;

    private string $promoShowed;

    private int $saveCc;

    private int $notUseBonuses;

    private int $frd;

    private int $frdPay;

    private int $frdTime;

    private int $frdCorr;

    private int $timeZoneId;

    private int $theme;

    private string $lastChange;

    private string $tempPassTime;

    private int $highVolumeFlag;

    private int $superVip;

    private string $suspendTime;


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