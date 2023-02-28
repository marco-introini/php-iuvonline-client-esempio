<?php

require __DIR__.'/../vendor/autoload.php';

$dotenv = Dotenv\Dotenv::createImmutable(__DIR__.'/..');
$dotenv->load();

$options = [
    'soap_version' => SOAP_1_1,
    'trace' => true,
    'keep_alive' => false,
    'connection_timeout' => 10000,
    'cache_wsdl' => WSDL_CACHE_NONE,
    'exceptions' => 0,
    'authentication' => SOAP_AUTHENTICATION_DIGEST,
    'stream_context' => stream_context_create([
        'ssl' => [
            'crypto_method' => STREAM_CRYPTO_METHOD_TLSv1_2_CLIENT,
            'cafile' => $_ENV['CA_CERT'],
            'local_cert' => $_ENV['AUTH_CERT'],
            'local_pk' => $_ENV['AUTH_KEY'],
            'verify_peer' => true,
            'verify_peername' => false,
            'allow_self_signed' => true
        ]
    ])
];

$soapClient = new SoapClient('../wsdl/wsdl/IUVOnlineService-scrittura-wsdl-1.3.wsdl', $options);
$xml = file_get_contents('../xml/create.xml');
$args = array();
if (!empty($xml)) {
    $args = array(new \SoapVar($xml, XSD_ANYXML));
}

$response = $soapClient->__soapCall('IUVOnlineCreate', $args);

echo $response;
