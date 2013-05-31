#!/usr/bin/env perl
use strict;
use warnings;
use utf8;
use 5.010000;
use autodie;

sub fib {
    my $n = shift;
    $n < 2 ? 1 : fib($n-1)+fib($n-2)
}

my $n = @ARGV > 0 ? $ARGV[0] : 13;
printf "fib(%d)=%d\n", $n, fib($n);

