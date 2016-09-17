<?php

$file = 'betclic.xml';

ini_set('pcre.backtrack_limit', 1024 * 1024 * 1024);
print 1;

$t = file_get_contents($file);
if(!preg_match_all('#<(([a-z\-_\.:]*)[^>]*)/?>#isu', $t, $f, PREG_SET_ORDER)) {
	die('Can\'t find tags');
}
print 2;

$stat = array();
$n = 0;
for($i=0; $i<count($f); $i++) {
	$v = &$f[$i];
	print '.';
	if($v[1][0] == '/') { // skip closing tags
		continue;
	}
	@$stat[$v[2]]['cnt']++;

	if(!preg_match_all('#([a-z\-_\.:]*)\=#isu', $v[1], $m, PREG_SET_ORDER)) {
		print_r($v);
		continue;
	}
	foreach($m as &$b) {
		@$stat[$v[2]]['tag'][$b[1]]++; // 'attr' of couse
	}
}
print 3;

ob_start();
print_r($stat);
file_put_contents($file.'-lookup.txt', ob_get_clean());
print 4;

?>