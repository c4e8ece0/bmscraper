<?php

$start = 'betfred-templates.xml';

$listxml = file_get_contents($start);
if(!$listxml) {
	die('Can`t fetch $start');
}

preg_match_all('#<template.*?name="([^"]*)"#isu', $listxml, $f);
if(empty($f[1])) {
	die('XML-links not found');
}

foreach($f[1] as $url) {
	print '.';
	$url = 'http://xml.betfred.com/'.$url.'.xml';

	if(file_exists('betfred/'.md5($url).'.xml')) {
		print '#';
		continue;
	}

	$t = file_get_contents($url);
	sleep(1);
	if($t) {
		file_put_contents('betfred/'.md5($url).'.xml', $t);
		file_put_contents('betfred/'.md5($url).'.url', $url);
		print '+';
	} else {
		print '-';
	}
}


?>