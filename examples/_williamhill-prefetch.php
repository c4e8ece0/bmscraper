<?php

$start = 'http://pricefeeds.williamhill.com/bet/en-gb?action=GoPriceFeed';

$listxml = file_get_contents($start);
if(!$listxml) {
	die('Can`t fetch $start');
}

preg_match_all('#<a href="([^"]*)">XML</a>#isu', $listxml, $f);
if(empty($f[1])) {
	die('XML-links not found');
}

foreach($f[1] as $url) {
	print '.';

	if(file_exists('williamhill/'.md5($url).'.xml')) {
		print '#';
		continue;
	}

	$t = file_get_contents($url);
	sleep(1);
	if($t) {
		file_put_contents('williamhill/'.md5($url).'.xml', $t);
		file_put_contents('williamhill/'.md5($url).'.url', $url);
		print '+';
	} else {
		print '-';
	}
}


?>