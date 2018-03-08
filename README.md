> # Source site `https://couponsfromchina.com` might be down. That is why parser is not able to fetch requred information.

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
 - Removes `bin` directory in the root of repository.
 - Creates binary with name `lsgbc` by default for host platform with in `bin` directory.

#### Without `make` utility

If you get stuck with `make` command, e.g. on `windows` platform, you can
perform all steps manually:

 - Build executable file for your platform from root directory of repository.
    ```
    $ go build
    ```
 - Check created binary.
    ```
    $ ./lsgbc -h
    Usage of ./dist/lsgbc-linux-amd64:
      -B, --best=false: show only best deals
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
`Eagle Eye` flashlights with price less than `20 ye`. To sort out names you can
specify begging of name like `-n jet`, this is equivalent of `-n "Jet*"`.

> __Note:__ by default rows are sorted by ascending the `PRICE`.

This can be achieved with following command:

```
$ ./lsgbc-linux-amd64 -M 20 -n jet,eagle

  NU  |                  NAME                  | PRICE, $ | DISCOUNT, % | LOWEST, $  
+-----+----------------------------------------+----------+-------------+-----------+
  155 | Jetbeam JET-u Flashlight               |     10.9 |           - |       7.0  
  108 | Eagle Eye X6 HOST Flashlight           |     14.1 |        10.0 |      11.0  
  156 | JETBeam JET-UV Flashlight              |     14.9 |           - |      11.6  
  152 | Jetbeam JET-I MK Flashlight            |     15.0 |           - |      10.0  
  150 | JETBeam i4 PRO Battery Charger EU Plug |     15.9 |        10.0 |      11.7  
  102 | Eagle Eye X2R 6000-6500K Flashlight    |     16.5 |        10.0 |      14.0  
  103 | Eagle Eye X2R NW Flashlight            |     16.9 |        10.0 |      12.5  
+-----+----------------------------------------+----------+-------------+-----------+
                                                               ITEMS    |     7      
                                                          +-------------+-----------+

```

There are following fields:

 - `NO`: number of product in the corresponding table.
 - `NAME`: product name from table.
 - `PRICE`: product price with applied coupon.
 - `DISCOUNT`: discount in percents comparing with regular price without coupon.
 - `LOWEST`: lowest price for this product during monitoring.

### Compact representation

When you get familiar with column names, you might want to use compact mode
`-C/--compact` on small terminals.

```
$ ./lsgbc-linux-amd64 -M 20 -n Jet,Eagle --compact

   #  |                   N                    | P, $ | D, % | L, $  
+-----+----------------------------------------+------+------+------+
  155 | Jetbeam JET-u Flashlight               | 10.9 |    - |  7.0  
  108 | Eagle Eye X6 HOST Flashlight           | 14.1 | 10.0 | 11.0  
  156 | JETBeam JET-UV Flashlight              | 14.9 |    - | 11.6  
  152 | Jetbeam JET-I MK Flashlight            | 15.0 |    - | 10.0  
  150 | JETBeam i4 PRO Battery Charger EU Plug | 15.9 | 10.0 | 11.7  
  102 | Eagle Eye X2R 6000-6500K Flashlight    | 16.5 | 10.0 | 14.0  
  103 | Eagle Eye X2R NW Flashlight            | 16.9 | 10.0 | 12.5  
+-----+----------------------------------------+------+------+------+
                                                                7    
                                                             +------+
```

### Flash sale and history

Attentive users might have noticed that some items do not have corresponding
`DISCOUNT` percent or `LOWEST` price values. Instead, this value might be
replaced with `-` char.

Dash `-` character as value has following meanings:

- `DISCOUNT`: you can buy this item without using of coupon (flash sale).
- `LOWEST`: there is no information about the lowest price for this item.

### The best deal

`lsgbc` allows to distinguish "best" deals with option `-B/--best`. This filter
shows only items with current `PRICE` equal or less than `LOWEST*1.1`.

For example:

```
$ ./lsgbc-linux-amd64 --max-price=15 --names=convoy --best

  NU |              NAME               | PRICE, $ | DISCOUNT, % | LOWEST, $  
+----+---------------------------------+----------+-------------+-----------+
  70 | Convoy S2+ CW Flashlight [GW4]  |     12.0 |        11.8 |      11.0  
  62 | Convoy S2 V2-1A Flashlight      |     12.3 |           - |      12.3  
  60 | Convoy S2 U6-3A Grey Flashlight |     12.7 |           - |      12.0  
  61 | Convoy S2 U6-4B Flashlight      |     12.8 |           - |      12.6  
+----+---------------------------------+----------+-------------+-----------+
                                                       ITEMS    |     4      
                                                  +-------------+-----------+
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
