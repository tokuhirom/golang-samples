use strict;
use warnings;
use utf8;
use Test::More;

is(`go run samples/coreutils/cat.go README.md`, slurp('README.md'));
is(`go run samples/coreutils/cat.go README.md README.md`, slurp('README.md')x2);
is(`go run samples/coreutils/cat.go < README.md`, slurp('README.md'));

done_testing;

sub slurp {
    my $fname = shift;
    open my $fh, '<', $fname
        or Carp::croak("Can't open '$fname' for reading: '$!'");
    scalar(do { local $/; <$fh> })
}
