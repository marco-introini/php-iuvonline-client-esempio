# Esempio di client PHP per accedere a IUVOnline BPS

### Installazione wsdl

Inserire il pacchetto wsd e xsd inviato dalla banca nella directory wsdl. Il risultato è il seguente:

![wsdl.png](img%2Fwsdl.png)

Modificare il wsdl immettendo l'endpoint inviato

### Installazione certificati

Copiare i certificati inviati nella directory cert. Il risultato sarà il seguente:

![certificati.png](img%2Fcertificati.png)

### Environment

Rinominare ed eventualmente modificare il file .env.example in .env

### XML della chiamata

Create un file xml valido per la chiamata nella directory xml

### Composer

Eseguire `composer install` per installare le dipendenze

### Chiamata Effettiva

Ora è possibile eseguire il file chiamata.php dentro la directory src

# Certificati CA

I certificati della CA e Intermediate CA sono presenti nella directory public_CA
