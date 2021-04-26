# ventiPi

A ventilation controller for the Raspberry Pi or any other device on which you can install and run Node-Red.
Still under development and without real function.

## Required hardware:

* A Raspberry Pi or other Hardware running Node-Red
* 2x [NodeMCU Lolin V3 Module ESP8266](https://www.amazon.de/gp/product/B074Q2WM1Y/ref=ppx_yo_dt_b_asin_title_o06_s00?ie=UTF8&psc=1)
* 2x [BME280 Sensor](https://www.amazon.de/gp/product/B07KY8WY4M/ref=ppx_yo_dt_b_asin_title_o00_s00?ie=UTF8&psc=1)
* 2x [HS100 Smart Plug](https://www.amazon.de/dp/B07TFGFVBW/ref=cm_sw_em_r_mt_dp_6KTB1F7CBCXTGHE7KKKK?_encoding=UTF8&psc=1 )
* [Ventilator](https://www.amazon.de/gp/product/B00V4VDA9C/ref=ppx_yo_dt_b_asin_title_o01_s00?ie=UTF8&psc=1)
* [Dehumidifier](https://www.amazon.de/gp/product/B07MSL8YN7/ref=ppx_yo_dt_b_search_asin_title?ie=UTF8&psc=1)

## Setup:

### Prepare ESP8266 Devices

#### Solder Sensor to ESP8266

Connect BME280 sensor to ESP8266, check `pinout` to see your layout.

```
   ESP8266 v3

 (A0 )    (D0 )     
 (GND)    (D1 )    
 (VU )    (D2 )   
 (S3 )    (D3 )
 (S2 )    (D4 )
 (S1 )    (3V3)
 (SC )    (GND)
 (S0 )    (D5 )
 (SK )    (D6 )
 (GND)    (D7 )               BME280            
 (3V3)    (D8 )
 (EN )    (RX )   ------->   (SCL)
 (RST)    (TX )   ------->   (SDA)
 (GND)    (GND)   ------->   (GND)
 (VIN)    (3V3)   ------->   (VIN)
    | USB  |
    | Port |
```


#### Tasmotize

Download Tasmotizer from [here](https://github.com/tasmota/tasmotizer).

##### Tasmotize your ESP8266 and configure WiFi

First select COM-Port and tasmota-sensors.bin, then click Tasmotize!  
After this finished, click on "Send config" and configure your WiFi-Settings. Leave everything else default.  
Click on "Get IP" to show the IP-Adress of the ESP8266 and paste this into your Brwoser.  
  
![Tasmotize and Configure](/tasmota/prepare/1.png "Tasmotize and Configure")


#### Configure the ESP8266 to post sensor data to MQTT

Click on Configuration -> Configure Module -> Select Module Type "Generic (18)" and click save  

![Configuration](/tasmota/prepare/2.png "Configuration")
![Configure Module](/tasmota/prepare/3.png "Configure Module")
![Select Module Type](/tasmota/prepare/4.png "Select Module Type")


Wait for Reboot, then Click Configuration -> Configure Module  
Set TX to I²C SCL  
Set RX to I²C SDA  
Click Save and wait for Reboot  
  
![Configure Module](/tasmota/prepare/5.png "Configure Module")


Sensor-Data is now shown on ESP8266-Homepage  

![Sensor-Data](/tasmota/prepare/6.png "Sensor-Data")


Configure Hostname on Configuration -> Configure WiFi -> Hostname, click Save  

![Hostname](/tasmota/prepare/7.png "Hostname")


Configure MQTT-Parameters on Configuration -> Configure MQTT  

Set Host, Port, Client, User to your needs  

Set topic an full topic to tasmota-innen for the inside sensor, tasmota-aussen for the outside sensor. Or set the names that you prefer and change it in the node-RED flow.  

![MQTT-Parameters](/tasmota/prepare/8.png "MQTT-Parameters")


Set the telemetry period to the value you prefer.  

![Telemetry period](/tasmota/prepare/9.png "Telemetry period")

#### Configure Sensor Offsets for Temp and Humidity
Like described [here](https://tasmota.github.io/docs/Commands/)

Open Webinterface of ESP8266 and Click on Console, then enter Command an hit Enter.

> TempOffset	-12.6..12.6 = Set calibraton offset value for reported temperature telemetry  
> This setting affects all temperature sensors on the device.  

```TempOffset -3 ```  

> HumOffset	-10.0..10.0 = Set calibraton offset value for reported humidity telemetry  
> This setting affects all humidity sensors on the device.  

```HumOffset 5 ```  

#### Install mosquitto Broker

```sudo apt update && sudo apt install mosquitto```

#### Install node-RED

Install node-Red like described [here](https://nodered.org/docs/getting-started/raspberrypi).