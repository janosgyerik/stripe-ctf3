#!/bin/sh

test "$1" && f="$1" || f=/usr/share/dict/words
f="$f" perl -pe '
BEGIN { 
  open(my $f, $ENV{f});
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
}
s/([^ \n]+)/&ff($+)/eg;
'
