# ventiPi

## German:
Eine Lüftungssteuerung für den Raspberry Pi oder jedes andere Gerät auf dem man Node-Red installieren und starten kann.
Ursprünglich wollte ich die Steuerung in Go schreiben, nun werde ich es wohl, aufgrund der guten Vorarbeit von SirReal auf [Youtube](https://www.youtube.com/watch?v=GmuldtHoURc&list=PLgYS2FpH2f4ruZqeL_5Q-74yXQD8xVN4L) in Node-Red umsetzen.

## English:
A ventilation controller for the Raspberry Pi or any other device on which you can install and run Node-Red.
Originally I wanted to write the control in Go, now I will probably implement it in Node-Red, due to the good preliminary work of SirReal on [Youtube](https://www.youtube.com/watch?v=GmuldtHoURc&list=PLgYS2FpH2f4ruZqeL_5Q-74yXQD8xVN4L).

## Required hardware:

* A Raspberry Pi or other Hardware running Node-Red
* [NodeMCU Lolin V3 Module ESP8266](https://www.amazon.de/gp/product/B074Q2WM1Y/ref=ppx_yo_dt_b_asin_title_o06_s00?ie=UTF8&psc=1)
* [SHT31 Sensor](https://www.amazon.de/gp/product/B07YQWX6BP/ref=ppx_yo_dt_b_asin_title_o05_s01?ie=UTF8&psc=1)


## Setup:

### Prepare ESP8266 Devices

#### Solder Sensor to ESP8266

Connect SHT31 sensor to ESP8266, check `pinout` to see your layout.

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
 (GND)    (D7 )               SHT31            
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


#### Install node-RED

Install node-Red like described [here](https://nodered.org/docs/getting-started/raspberrypi).