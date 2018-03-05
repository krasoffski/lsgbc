> # Source site `https://couponsfromchina.com` is down. That is why parser is not able to fetch requred information.

# lsgbc - List GearBest Coupons
[![Build Status](https://travis-ci.org/krasoffski/lsgbc.svg?branch=master)](https://travis-ci.org/krasoffski/lsgbc)
[![Code Climate](https://codeclimate.com/github/krasoffski/lsgbc/badges/gpa.svg)](https://codeclimate.com/github/krasoffski/lsgbc)

## Objectivity
In short words, `lsgbs` is a command line interface for
https://couponsfromchina.com/ which allows you to sort, include or exclude
different types of items for the best deal look up.

## Installation

There are compiled binaries for `amd64` architecture for linux, windows and
darvin platforms. Please, find these files in the release menu of the project.

For creating binaries you need `Git` and `Golang` compiler with version equal
or bigger than `1.8` installed. Please, make sure you have properly configured
`GOPATH` and system path variables:
 - Add `$GOROOT/bin` to system path.
 - Add `$GOPATH/bin` to system path.

### Manual installation

Before all, please clone this repository with command:

```bash
$ git clone https://github.com/krasoffski/lsgbc
# git clone git@github.com:krasoffski/lsgbc.git
```

And change working directory to cloned repository: `cd lsgbc`

#### Using `make all` target

Installation with `make all` target performs following steps (requires `make`):
 - Download and install [`dep`](https://github.com/golang/dep) package manager.
 - Download dependencies with `dep` command to `vendor` directory.
 - Create binaries for windows, linux and darvin in the `dist` directory.

#### Without `make` utility

If you get stuck with `make` command, e.g. on `windows` platform, you can
perform all steps manually:

 - Install `dep` package manager (see note about system path) if it not present.
    ```
    $ go get -u github.com/golang/dep/cmd/dep
    ```
 - Download required dependencies with `dep` command.
    ```
    $ dep ensure
    ```
 - Build executable file for your platform.
    ```
    $ go build
    ```
 - Check created binary.
    ```
    $ ./dist/lsgbc-linux-amd64 -h
    Usage of ./dist/lsgbc-linux-amd64:
      -B, --best=false: show only best deals
      -c, --categories="*": comma separated list of categories (case sensitive), e.g. 'aa,b*,cc'
      -C, --compact=false: use compact table representation
      -F, --flash-sale=false: show only flash sale deals
      -l, --list="flashlight": used coupons list, one from: 3d,...,xiaomi
      -M, --max-price=1000: maximum discount price
      -m, --min-price=0: minimal discount price
      -n, --names="*": comma separated list of names (case sensitive), e.g. 'xx,y*,zz'
      -S, --sort-by="price": sort table by column, 'price' or 'discount'
      -V, --version=false: show version and exit
    ```

## Usage

For example, you would like to know current price with coupon for `Jetbeam` and
`Eagle Eye` flashlights with price less than `20 ye`. To sort out name/categories
you can specify begging of names, categories like `-n jet`, this is equivalent
of `-n "Jet*"`.

> __Note:__ by default rows are sorted by ascending the `PRICE`.

This can be achieved with following command:

```
$ ./lsgbc-linux-amd64 -M 20 -n jet,eagle

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

There are following fields:

 - `NO`: number of product in the corresponding table.
 - `NAME`: product name from table.
 - `PRICE`: product price with applied coupon.
 - `DISCOUNT`: discount in percents comparing with regular price without coupon.
 - `LOWEST`: lowest price for this product during monitoring.
 - `CATEGORY`: product category from URL path.

### Compact representation

When you get familiar with column names and categories, you might want to use
compact mode `-C/--compact` on small terminals.

```
$ ./lsgbc-linux-amd64 -M 20 -n Jet,Eagle --compact

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
`DISCOUNT` percent or `LOWEST` price values. Instead, this value might be
replaced with `-` char.

Dash `-` character as value has following meanings:

- `DISCOUNT`: you can buy this item without the use of coupon (flash sale).
- `LOWEST`: there is no information about the lowest price for this item.

### The best deal

`lsgbc` allows to distinguish "best" deals with option `-B/--best`. This filter
shows only items with current `PRICE` equal or less than `LOWEST*1.1`.

For example:

```
$ ./lsgbc-linux-amd64 --max-price=15 --categories=led --best

  NU  |                  NAME                   | PRICE, $ | DISCOUNT, % | LOWEST, $ |    CATEGORY
+-----+-----------------------------------------+----------+-------------+-----------+-----------------+
  384 | Zanflare F6S Flashlight                 |      6.0 |        70.0 |       6.0 | led-flashlights
  212 | Lumintop IYP365 CW Flashlight           |     10.0 |        20.7 |      10.0 | led-flashlights
  374 | YWXLight 5000 lm Headlamp               |     10.0 |        23.1 |      10.0 | led-flashlights
  216 | Lumintop Tool Nichia 219BT Flashlight   |     10.0 |        47.4 |       9.1 | led-flashlights
  213 | Lumintop IYP365 NW Flashlight           |     12.0 |           - |      12.0 | led-flashlights
   79 | Convoy S2+ 365nm Nichia UV Flashlight   |     13.0 |        35.0 |      12.0 | led-flashlights
   74 | Convoy S2 V2-1A Flashlight              |     13.0 |           - |      12.6 | led-flashlights
+-----+-----------------------------------------+----------+-------------+-----------+-----------------+
                                                                             ITEMS   |        7
                                                                         +-----------+-----------------+
```
This table contains only deals for:
 - Maximum `PRICE` is `$15`.
 - `CATEGORY` equals to `led-flashlights`.
 - `PRICE` is around of `LOWEST`.
 - Items sorted by ascending the `PRICE`.

E.g. the `PRICE` of `Lumintop Tool Nichia 219BT Flashlight` is `$10` and this
less than `LOWEST` price multiply by `1.1` (10.0 < 10.01=9.1*1.1).

### Sorting items

You can sort items by `PRICE` or by `DISCOUNT`. This can be done using option
`-S/--sort-by` with value `discount` or `d` for shortness to sort by decreasing
of `DISCOUNT` percents.

```
$ ./dist/lsgbc-linux-amd64 -C -M 30 -c led -n Lumintop -S d

   #  |                       N                        | P, $ | D, % | L, $
+-----+------------------------------------------------+------+------+------+
  184 | Lumintop Copper Tool AAA XP-G2 R5 Flashlight   | 21.0 | 30.0 | 20.0
  188 | Lumintop IYP365 CW Flashlight                  | 10.0 | 20.7 | 10.0
  183 | Lumintop Copper Tool AAA Nichia 219 Flashlight | 20.0 |    - | 20.0
  189 | Lumintop IYP365 NW Flashlight                  | 12.0 |    - | 12.0
  197 | Lumintop Tool LED Keychain Flashlight          | 10.0 |    - |  9.0
  198 | Lumintop Tool Nichia 219BT Flashlight          | 10.0 |    - |  9.1
+-----+------------------------------------------------+------+------+------+
                                                                        6
                                                                     +------+
```

Items with flash sale has empty discount field. As result, discount for these
items is replaced with `-` and shown at the end of the table (zero discount).
