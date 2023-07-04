<?php
/* This is a simple PHP example to host your own Amazon Alexa Skill written in PHP.
In my Case it connects to my smarthome Raspberry pi Cat Feeder with two intents;
1: Dispense Food to the cats.
2: When did the Feeder last time feed the cats? Return a spoken time / date
This Script contains neccessary calls and security to give you a easy to use DIY example.

v2016.12.29    
Details in my Blogpost:  https://solariz.de/de/amazon-echo-alexa-meets-catfeeder.htm
*/

//bu kod puplic key i cikarir 
//openssl x509 -pubkey -noout -in private.pem >> public.key 


// sudo openssl dgst -sha1 -sign private.pem -out /tmp/selman.sha1 selman.sha1
// sudo openssl base64 -in /tmp/selman.sha1 -out signature.sha1
// rm /tmp/selman.sha1


// #verify 

// sudo openssl base64 -d -in signature.sha1 -out /tmp/selman.sha1
// sudo openssl dgst -sha1 -verify public.key -signature /tmp/selman.sha1 selman.sha1
// rm /tmp/selman.sha1


header('Cache-Control: no-cache, must-revalidate');
header('Expires: Mon, 26 Jul 1997 05:00:00 GMT');

// SETUP / CONFIG

$SETUP = array(
	'SkillName' => "restaurant",
	'SkillVersion' => '1.0',
	'ApplicationID' => 'amzn1.ask.skill.72d8ce35-6532-481f-aecb-b7149015f763', // From your ALEXA developer console like: 'amzn1.ask.skill.45c11234-123a-1234-ffaa-1234567890a'
	'CheckSignatureChain' => false, // make sure the request is a true amazonaws api call
	'ReqValidTime' => 60, // Time in Seconds a request is valid
	'AWSaccount' => 'amzn1.ask.account.AMASNF3KSDMAQSKOF7U4UGDKWHFGWCTU562KUYGJZEI55LHUIZ44QZI3CFTJHNFVQGJTXLQ735FLVXU7UML4O5MJCED4AKG7XGNIVNQTV4V3GDTC6W2RAMYKZNSZ6TCGJNEM7ZNRFIRXNWAPYCJIK6UQVP4LQCT4JZAC4VZGLNGKX3WF2IVDMCOXT2XCKQOAF42X7IAHZWHTQ7TJ5HDFYR37OXMQXMWOKVW3JAHD7BXA', //If this is != empty the specified session->user->userId is required. This is usefull for account bound private only skills
	'validIP' => array(
		"72.21.217.",
		"54.240.197."
	) , // Limit allowed requests to specified IPv4, set to FALSE to disable the check.
	'LC_TIME' => "de_DE"

	// We use german Echo so we want our date output to be german

);
setlocale(LC_TIME, $SETUP['LC_TIME']);

// Getting Input

$rawJSON = file_get_contents('php://input');
$EchoReqObj = json_decode($rawJSON);

if (is_object($EchoReqObj) === false) ThrowRequestError();
$RequestType = $EchoReqObj->request->type;

// Check if Amazon is the Origin

// if (is_array($SETUP['validIP'])){
// 	$isAllowedHost = false;
// 	foreach($SETUP['validIP'] as $ip)
// 		{
// 		if (stristr($_SERVER['REMOTE_ADDR'], $ip))
// 			{
// 			$isAllowedHost = true;
// 			break;
// 			}
// 		}
// if ($isAllowedHost == false) ThrowRequestError(403, "Forbidden, your Host is not allowed to make this request!");
// 	unset($isAllowedHost);
// }


// Check if correct requestId
if (strtolower($EchoReqObj->session->application->applicationId) != strtolower($SETUP['ApplicationID']) || empty($EchoReqObj->session->application->applicationId))
	{
	ThrowRequestError(401, "Forbidden, unkown Application ID!");
	}

// Check SSL Signature Chain







	// Check Valid Time

if (time() - strtotime($EchoReqObj->request->timestamp) > $SETUP['ReqValidTime']) ThrowRequestError(408, "Request Timeout! Request timestamp is to old.");

// Check AWS Account bound, if this is set only a specific aws account can run the skill

if (!empty($SETUP['AWSaccount']))
	{
	if (empty($EchoReqObj->session->user->userId) || $EchoReqObj->session->user->userId != $SETUP['AWSaccount'])
		{
		ThrowRequestError(403, "Forbidden! Access is limited to one configured AWS Account.");
		}
}

$JsonOut = GetJsonMessageResponse($RequestType, $EchoReqObj);
header('Content-Type: application/json');
header("Content-length: " . strlen($JsonOut));
echo $JsonOut;
exit();

// -----------------------------------------------------------------------------------------//
//					     functions
// -----------------------------------------------------------------------------------------//
// This function returns a json blob for output

function GetJsonMessageResponse($RequestMessageType, $EchoReqObj)
	{
	GLOBAL $SETUP;
	$RequestId = $EchoReqObj->request->requestId;
	$ReturnValue = "";
	if ($RequestMessageType == "LaunchRequest")
		{
		$return_defaults = array(
			'version' => $SETUP['SkillVersion'],
			'sessionAttributes' => array(
				'countActionList' => array(
					'read' => true,
					'category' => true
				)
			) ,
			'response' => array(
				'outputSpeech' => array(
					'type' => "PlainText",
					'text' => "Willkommen beim CatFeeder Beispiel"
				) ,
				'card' => array(
					'type' => "Simple",
					'title' => "CatFeeder",
					'content' => "Test Content"
				) ,
				'reprompt' => array(
					'outputSpeech' => array(
						'type' => "PlainText",
						'text' => "Kann ich dir noch weiter behilflich sein?"
					)
				)
			) ,
			'shouldEndSession' => true
		);
		$ReturnValue = json_encode($return_defaults);
		}
	elseif ($RequestMessageType == "SessionEndedRequest")
		{
		$ReturnValue = json_encode(array(
			'type' => "SessionEndedRequest",
			'requestId' => $RequestId,
			'timestamp' => date("c") ,
			'reason' => "USER_INITIATED"
		));
		}
	elseif ($RequestMessageType == "IntentRequest")
		{
		if ($EchoReqObj->request->intent->name == "CatFeederFeed") // Alexa Intent name
			{

			// do what ever your intent should do here. In my Case I call home to my raspberry pi, see function comment for more info.

			getRequestPayload(array(
				'action' => "feed",
				'size' => 1
			));
			$SpeakPhrase = "OK";
			}
		elseif ($EchoReqObj->request->intent->name == "CatFeederLast") // 2nd Alexa Intent name
			{

			// do what ever your intent should do here. In my Case I call home to my raspberry pi, see function comment for more info.

			$last_feed = getRequestPayload(array(
				'action' => "LastFeed"
			));
			if (!$last_feed || $last_feed <= 1000000000)
				{ // Should return a linux timestamp, check if plausible, else Error.
				$SpeakPhrase = "Ich kann diese Information momentan leider nicht ermitteln.";
				}
			elseif (time() - $last_feed < 900) $SpeakPhrase = "Die Katzen wurden gerade eben gefÃ¼ttert.";
			elseif (time() - $last_feed < 3600) $SpeakPhrase = "Die Katzen wurden vor weniger als einer Stunde gefÃ¼ttert.";
			elseif (time() - $last_feed < 7200) $SpeakPhrase = "Die Katzen wurden vor eins bis zwei Stunden gefÃ¼ttert.";
			  else
				{ // More human readable Date formating:
				if (strftime("%e") == strftime("%e", $last_feed))
					{
					$day = "heute";
					}
				elseif (intval(strftime("%e", $last_feed)) == intval(strftime("%e") - 1))
					{
					$day = "gestern";
					}
				  else
					{
					$day = strftime("%e. %B", $last_feed);
					}

				$SpeakPhrase = "Die letzte FÃ¼tterung war " . $day . " um " . strftime("%H:%M", $last_feed);
				}
			}

		$ReturnValue = json_encode(array(
			'version' => $SETUP['SkillVersion'],
			'sessionAttributes' => array(
				'countActionList' => array(
					'read' => true,
					'category' => true
				)
			) ,
			'response' => array(
				'outputSpeech' => array(
					'type' => "PlainText",
					'text' => $SpeakPhrase
				) ,
				'card' => array(
					'type' => "Simple",
					'title' => "CatFeeder",
					'content' => $SpeakPhrase
				)
			) ,
			'shouldEndSession' => true
		));
		}
	  else
		{
		ThrowRequestError();
		}

	return $ReturnValue;
	} // end function GetJsonMessageResponse

function ThrowRequestError($code = 400, $msg = 'Bad Request')
	{
	GLOBAL $SETUP;
	http_response_code($code);
	echo "Error " . $code . "<br />\n" . $msg;
	error_log("alexa/" . $SETUP['SkillName'] . ":\t" . $msg, 0);
	exit();
	}

function getRequestPayload($payload)
	{
	/* this is just a custom function to get my connection to my home device.
	In this example it's a raspberry pi based CatFeeder listening to POST requests.
	It is using a let's encrypt SSL cert and a basic HTTP Auth.
	Use it as a example to DIY here: */
	$username = "basicauthuser";
	$password = "basicauthpasswd";
	$host = "https://myawesome.raspberry.pi.at.home/catfeed_post.php";
	$process = curl_init($host);
	curl_setopt($process, CURLOPT_SSL_VERIFYHOST, 2);
	curl_setopt($process, CURLOPT_SSL_VERIFYPEER, 1);
	curl_setopt($process, CURLOPT_USE_SSL, CURLUSESSL_ALL);
	curl_setopt($process, CURLOPT_SSL_VERIFYSTATUS, 1);
	curl_setopt($process, CURLOPT_HEADER, FALSE);
	curl_setopt($process, CURLOPT_USERPWD, $username . ":" . $password);
	curl_setopt($process, CURLOPT_TIMEOUT, 8);
	curl_setopt($process, CURLOPT_POST, 1);
	curl_setopt($process, CURLOPT_VERBOSE, FALSE);
	curl_setopt($process, CURLOPT_POSTFIELDS, http_build_query($payload));
	curl_setopt($process, CURLOPT_RETURNTRANSFER, TRUE);
	$return = curl_exec($process);
	curl_close($process);
	return $return;
	}




	if ($SETUP['CheckSignatureChain'] == true){
		if (preg_match("/https:\/\/s3.amazonaws.com(\:443)?\/echo.api\/*/i", $_SERVER['HTTP_SIGNATURECERTCHAINURL']) == false)
			{
			ThrowRequestError(403, "Forbidden, unkown SSL Chain Origin!");
			}
	
		// PEM Certificate signing Check
		// First we try to cache the pem file locally
	
		$local_pem_hash_file = sys_get_temp_dir() . '/' . hash("sha256", $_SERVER['HTTP_SIGNATURECERTCHAINURL']) . ".pem";
		print_r($local_pem_hash_file);
		if (!file_exists($local_pem_hash_file)){
			file_put_contents($local_pem_hash_file, file_get_contents($_SERVER['HTTP_SIGNATURECERTCHAINURL']));
		}
		$local_pem = file_get_contents($local_pem_hash_file);
		if (openssl_verify($rawJSON, base64_decode($_SERVER['HTTP_SIGNATURE']) , $local_pem) !== 1){
			ThrowRequestError(403, "Forbidden, failed to verify SSL Signature!");
		}
	
		// Parse the Certificate for additional Checks
	
		$cert = openssl_x509_parse($local_pem);
		if (empty($cert)) ThrowRequestError(424, "Certificate parsing failed!");
	
		// SANs Check
	
		if (stristr($cert['extensions']['subjectAltName'], 'echo-api.amazon.com') != true) ThrowRequestError(403, "Forbidden! Certificate SANs Check failed!");
	
		// Check Certificate Valid Time
	
		if ($cert['validTo_time_t'] < time())
			{
			ThrowRequestError(403, "Forbidden! Certificate no longer Valid!");
	
			// Deleting locally cached file to fetch a new at next req
	
			if (file_exists($local_pem_hash_file)) unlink($local_pem_hash_file);
			}
	
		// Cleanup
	
		unset($local_pem_hash_file, $cert, $local_pem);
		die;
		}