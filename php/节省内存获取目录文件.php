<?php

function getDirectoryFiles($directory) {
  $files = glob("$directory/*");
  foreach ($files as $file) {
    yield $file;
  }
}

$generator = getDirectoryFiles('/path/to/directory');

foreach ($generator as $file) {
  echo $file . "\n";
}
