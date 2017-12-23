# lsgbc - List GearBest Coupons
[![Build Status](https://travis-ci.org/krasoffski/lsgbc.svg?branch=master)](https://travis-ci.org/krasoffski/lsgbc)

## Objectivity
In short words, `lsgbs` is a command line interface for
https://couponsfromchina.com/ which allows you to sort, include or exclude
different types of items for the best deal look up.

## Usage

For example, you would like to know current price with coupon for `Jetbeam` and
`Eagle Eye` flashlights with price less than `20 ye`. You can use `*` for
matching to avoid full name/category typing.

This can be achieved with following command:

```
$ ./lsgbc-linux-amd64 -M 20 -n 'Jet*,Eagle*'

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

> __Note:__ `-n/--names` and `-c/--categories` options are case sensitive.

There are following fields:

 - `NO`: number of product in the corresponding table.
 - `NAME`: product name from table.
 - `PRICE`: product price with applied coupon.
 - `DISCOUNT`: discount in percents comparing with regular price without coupon.
 - `LOWEST`: lowest price for this product during monitoring.
 - `CATEGORY`: product category from URL path.

### Compact representation

When you get familiar with column names and categories, you might want to use compact mode
`-C/--compact` on small terminals.

```
$ ./lsgbc-linux-amd64 -M 20 -n 'Jet*,Eagle*' --compact

   #  |                  N                  | P, $ | D, % | L, $
+-----+-------------------------------------+------+------+------+
  164 | Jetbeam JET-u Flashlight            | 11.0 |    - |  7.0
  111 | Eagle Eye X6 HOST Flashlight        | 13.0 | 13.4 | 11.0
  105 | Eagle Eye X2R 6000-6500K Flashlight | 14.0 | 22.5 |    -
  163 | Jetbeam JET-I MK Flashlight         | 14.0 |    - | 10.0
  103 | Eagle Eye X2R 1A Flashlight         | 15.9 | 12.0 | 15.0
  104 | Eagle Eye X2R 3C Flashlight         | 16.3 | 12.0 | 13.7
  106 | Eagle Eye X2R NW Flashlight         | 16.3 | 12.0 | 12.5
  107 | Eagle Eye X5R 3A Flashlight         | 20.0 | 34.8 | 20.0
+-----+-------------------------------------+------+------+------+
                                                             8
                                                          +------+
```

### Flash sale and history

Attentive users might have noticed that some items do not have corresponding
`DISOUNT` persent or `LOWEST` price values. Instead, this value might be
replaced with `-` char.

Dash `-` charaster as value has following meanings:

- `DISCOUNT`: you can buy this item without the use of coupon (flash sale).
- `LOWEST`: there is no information about the lowest price for this item.

### Best deal

`lsgbc` allows to distnguish "best" deasl with option `-B/--best`. This filter
shows only items with current `PRICE` equal or less than `LOWEST*1.1`.

For example:
```
$ ./dist/lsgbc-linux-amd64 --max-price=15 --categories="led*" --best

  NU  |                  NAME                   | PRICE, $ | DISCOUNT, % | LOWEST, $ |    CATEGORY
+-----+-----------------------------------------+----------+-------------+-----------+-----------------+
  351 | Ultrafire H19-1 Flashlight              |      5.4 |           - |       5.4 | led-flashlights
  384 | Zanflare F6S Flashlight                 |      6.0 |        70.0 |       6.0 | led-flashlights
  212 | Lumintop IYP365 CW Flashlight           |     10.0 |        20.7 |      10.0 | led-flashlights
  374 | YWXLight 5000 lm Headlamp               |     10.0 |        23.1 |      10.0 | led-flashlights
  216 | Lumintop Tool Nichia 219BT Flashlight   |     10.0 |        47.4 |       9.1 | led-flashlights
  213 | Lumintop IYP365 NW Flashlight           |     12.0 |           - |      12.0 | led-flashlights
   79 | Convoy S2+ 365nm Nichia UV Flashlight   |     13.0 |        35.0 |      12.0 | led-flashlights
   74 | Convoy S2 V2-1A Flashlight              |     13.0 |           - |      12.6 | led-flashlights
  123 | HaikeLite HT08 Flashlight               |     13.0 |        23.5 |      13.0 | led-flashlights
  259 | Nitecore TIP CRI Flashlight Blue        |     14.0 |        60.0 |      14.0 | led-flashlights
  263 | Nitecore TIP CRI Flashlight Silver      |     14.0 |        60.0 |      14.0 | led-flashlights
  271 | Nitecore TIP XP-G2 S3 Silver Flashlight |     15.0 |        50.0 |      15.0 | led-flashlights
  269 | Nitecore TIP XP-G2 S3 Flashlight Green  |     15.0 |        10.7 |      15.0 | led-flashlights
  268 | Nitecore TIP XP-G2 S3 Flashlight Golden |     15.0 |        10.7 |      15.0 | led-flashlights
  267 | Nitecore TIP XP-G2 S3 Flashlight Blue   |     15.0 |        10.7 |      15.0 | led-flashlights
  270 | Nitecore TIP XP-G2 S3 Flashlight Red    |     15.0 |        10.7 |      15.0 | led-flashlights
+-----+-----------------------------------------+----------+-------------+-----------+-----------------+
                                                                             ITEMS   |       16
                                                                         +-----------+-----------------+
```
This table contains only deals for `CATEGORY` equels to `led-flashlights` with
maximum `PRICE` is `$15` where `PRICE` is around of `LOWEST` seen price. E.g.
the `PRICE` of `Lumintop Tool Nichia 219BT Flashlight` is `$10` and this less
than `LOWEST` price multiply by `1.1` (10.0 < 10.01=9.1*1.1).
