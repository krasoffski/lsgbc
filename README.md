# lsgbc - List GearBest Coupons
[![Build Status](https://travis-ci.org/krasoffski/lsgbc.svg?branch=master)](https://travis-ci.org/krasoffski/lsgbc)

## Objectivity
In short words `lsgbs` is a command line interface for
https://couponsfromchina.com/ which allows you to sort, include or exclude
different types of items for best deal look up.

## Usage

For example, you would like to know current price with coupon for `Jetbeam`
flashlights with price less than `20 ye`. You can use `*` for matching to avoid
full name/category typing.

This can be achieved with following command:

```bash
$ ./lsgbc-linux-amd64 -max 20 -name 'Jet*'

  NO  |            NAME             | PRICE, $ | DISCOUNT, % | LOWEST, $ |    CATEGORY
+-----+-----------------------------+----------+-------------+-----------+-----------------+
  164 | Jetbeam JET-u Flashlight    |     11.0 |           - |       7.0 | led-flashlights
  163 | Jetbeam JET-I MK Flashlight |     16.3 |        12.0 |      10.0 | led-flashlights
+-----+-----------------------------+----------+-------------+-----------+-----------------+
                                                                 ITEMS   |        2
                                                             +-----------+-----------------+
```

_Note: `-name` and `-categories` options are case sensitive._

You can notice here such fields:

 - `NO` - number of product in the corresponding table.
 - `NAME` - product name from table.
 - `PRICE` - product price with applied coupon.
 - `DISCOUNT` - discount in percents comparing with regular price without coupon
 - `LOWEST` - lowest price for this product during monitoring
 - `CATEGORY` - product category from URL path
