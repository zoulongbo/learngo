<?php
/**
 * Created by Phpstorm.
 * User : zoulongbo
 * Date : 2020-03-06
 * Tome : 10:07
 */
return [
    "upload"                => [
        "file" => "上传文件无效"
    ],
    "email"                 => [
        "send_failed" => "邮件发送失败"
    ],
    "sms"                   => [
        "send_failed" => "短信发送失败"
    ],
    "orders"                => [
        "failed" => "下单失败"
    ],
    "logistics"             => [
        "order_status_not_allow" => "订单状态不允许物流采购",
        "data_insert_failed"     => "数据插入失败",
        "order_illegal"          => "存在非法订单",
        "expense_detail_empty"   => "费用明细还未生成，请稍后操作"
    ],
    "order"                 => [
        "user_not_auth"    => "用户未认证或未实名",
        "sign_not_auth"    => "用户未认证或未实名",
        "shop_can_not_pay" => "店铺未开通支付",
        "generate_failed"  => "下单失败"
    ],
    "oauth"                 => [
        "invalid_client"        => "无效的客户端",
        "authentication_failed" => "认证失败",
        "refresh_failed"        => "刷新认证失败",
        "invalid_code"          => "无效",
        "expired_code"          => "已过期",
    ],
    "cms"                   => [
        "sensitive_word_list_failed"       => "敏感词列表拉取失败",
        "sensitive_word_store_failed"      => "敏感词新增失败",
        "sensitive_word_destroy_failed"    => "敏感词删除失败",
        "sensitive_word_check_failed"      => "敏感词校验失败",
        "content_contains_sensitive_words" => "内容含有敏感词",
    ],
    "third"                 => [
        "user"             => [
            "register"   => [
                "accountFail1"              => "请输入您的邮箱或者手机！",
                "accountFail2"              => "密码过于简单！需要数字、字母、6-20位之间",
                "accountError1"             => "注册失败！",
                "accountError2"             => "该账号已注册，请更换联系方式。",
                "accountError3"             => "当前账户在1个月内存在注销账户的操作，不允许再次注册。",
                "accountSuccess"            => "注册成功！",
                "setMessageCodeFail1"       => "该账号已被注册，请更换联系方式。",
                "setMessageCodeFail2"       => "请选择正确的验证码接受方式！",
                "setMessageCodeError"       => "发送验证码失败！",
                "setMessageByEmail"         => "验证码发送失败！",
                "setMessageByPhoneFail"     => "错误的手机格式！",
                "checkAccountFail"          => "该账户不是登陆账号！",
                "checkAccountError"         => "判断是否为登陆账户失败！",
                "changeLoginAccountFail1"   => "新手机号不能为空！",
                "changeLoginAccountFail2"   => "您的手机号输入有误！",
                "changeLoginAccountFail3"   => "账户原本的手机号不能为空！",
                "changeLoginAccountFail4"   => "该手机已经被注册，请更换手机号！",
                "changeLoginAccountSuccess" => "变更登录账户成功！",
                "changeLoginAccountFail"    => "变更登录账户失败！",
            ],
            "changeInfo" => [
                "changePasswordFail3"                  => "判断用户联系方式失败！",
                "changePasswordFail4"                  => "旧密码错误！",
                "changePasswordSuccess"                => "修改登陆密码成功！",
                "setMessageCodeForChangePasswordFail1" => "找不到该账户！",
                "changeEmailFail1"                     => "userId不能为空！",
                "changeEmailFail2"                     => "查询不到该账户",
                "changeEmailFail3"                     => "该账户已被使用",
                "changeEmailSuccess"                   => "登录名修改成功！",
                "verifyIdCardFail"                     => "账号，身份证信息不匹配！",
                "changeUserInfo"                       => "不能所有信息都为空！",
                "insertEmail"                          => "修改用户信息成功",
                "checkIdCard"                          => "身份证信息匹配错误！",
                "getUserInfoByUserId"                  => "该用户信息不存在！",
            ],
        ],
        "regulation"       => [
            "subject" => [
                "empty" => "主体不能为空。",
            ],
            "code"    => [
                "empty" => "验证码不能为空。",
            ],
        ],
        "limit"            => [
            "isTotalRetryExceeded" => "发送次数已超限。",
            "isPerRetryPermit"     => [
                "canRepeat" => "验证码已发送, 有效期:10分钟。1分钟后您可以重新发送验证码。"
            ],
        ],
        "verificationcode" => [
            "verifycode" => [
                "invalid" => "验证码失效。",
                "error"   => "验证码错误。",
            ],
        ],
        "LimitExceeded"    => [
            "AppDailyLimit"                    => "业务短信日下发条数超过设定的上限",
            "DailyLimit"                       => "短信日下发条数超过设定的上限(国际/港澳台)",
            "DeliveryFrequencyLimit"           => "下发短信命中了频率限制策略",
            "PhoneNumberCountLimit"            => "调用短信发送API接口单次提交的手机号个数超过200个",
            "PhoneNumberDailyLimit"            => "单个手机号日下发短信条数超过设定的上限",
            "PhoneNumberOneHourLimit"          => "单个手机号1小时内下发短信条数超过设定的上限",
            "PhoneNumberSameContentDailyLimit" => "单个手机号下发相同内容超过设定的上限",
            "PhoneNumberThirtySecondLimit"     => "单个手机号30秒内下发短信条数超过设定的上限",
        ],
    ],
    "internal_server_error" => "服务器内部错误",
    "wx"        => [
        "open"  => [
            "config_not_exist"  => "配置不存在",
        ],
    ],
];
