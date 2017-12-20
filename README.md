# lsgbc - List GearBest Coupons
[![Build Status](https://travis-ci.org/krasoffski/lsgbc.svg?branch=master)](https://travis-ci.org/krasoffski/lsgbc)

## Objectivity
In short words `lsgbs` is a command line interface for
https://couponsfromchina.com/ which allows you to sort, include or exclude
different types of items for best deal look up.

## Usage

For example, you would like to know current price with coupon for `Jetbeam` and
`Eagle Eye` flashlights with price less than `20 ye`. You can use `*` for
matching to avoid full name/category typing.

This can be achieved with following command:

```bash
$ ./lsgbc-linux-amd64 -M 20 -n 'Jet*,Convoy*'

  NO  |                NAME                 | PRICE, $ | DISCOUNT, % | LOWEST, $ |    CATEGORY
+-----+-------------------------------------+----------+-------------+-----------+-----------------+
  164 | Jetbeam JET-u Flashlight            |     11.0 |           - |       7.0 | led-flashlights
  111 | Eagle Eye X6 HOST Flashlight        |     13.0 |        13.4 |      11.0 | led-flashlights
  105 | Eagle Eye X2R 6000-6500K Flashlight |     14.0 |        22.5 |         - | led-flashlights
  163 | Jetbeam JET-I MK Flashlight         |     14.0 |           - |      10.0 | led-flashlights
  103 | Eagle Eye X2R 1A Flashlight         |     15.9 |        12.0 |      15.0 | led-flashlights
  104 | Eagle Eye X2R 3C Flashlight         |     16.3 |        12.0 |      13.7 | led-flashlights
  106 | Eagle Eye X2R NW Flashlight         |     16.3 |        12.0 |      12.5 | led-flashlights
  107 | Eagle Eye X5R 3A Flashlight         |     20.0 |        34.8 |      20.0 | led-flashlights
+-----+-------------------------------------+----------+-------------+-----------+-----------------+
                                                                         ITEMS   |        8
                                                                     +-----------+-----------------+
```

_Note: `-n/--names` and `-c/--categories` options are case sensitive._

You can notice here such fields:

 - `NO` - number of product in the corresponding table.
 - `NAME` - product name from table.
 - `PRICE` - product price with applied coupon.
 - `DISCOUNT` - discount in percents comparing with regular price without coupon
 - `LOWEST` - lowest price for this product during monitoring
 - `CATEGORY` - product category from URL path
