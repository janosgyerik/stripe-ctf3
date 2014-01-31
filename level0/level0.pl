#!/usr/bin/perl

open(my $f, $ARGV[0]);
chomp(@w = <$f>);
@h{@w} = map(1, @w);
sub ff {
    $v = $_[0];
    if ($v =~ /\W/) {
        "<$v>"
    } else {
        if (exists($index{$v})) {
            $index{$v}
        } else {
            $index{$v} = $h{lc($v)} ? $v : "<$v>";
        }
    }
}
while (<>) {
    s/([^ \n]+)/&ff($+)/eg;
}
