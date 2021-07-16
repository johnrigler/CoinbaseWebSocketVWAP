# CoinbaseWebSocketVWAP
I use the term "Volume Weighted Average Price" a bit liberally. The basic formula
has been adapted for a Demo. The general formula is:

VWAP = (Typical Price * Volume)/Cumulative Volume 

Typical Price is defined as the Average of:
The High Price, The Low Price, and the Closing Price for that day (I use the current price instead)

I create an array of the volume of the most recent 200 "Buy" Matches for BTC-USD, ETH-USD, and BTC-ETH. By using
the modulus operator, I am able to alter this array after the 200th record in a pretty useful way.

This could be much better and should not be used for real market decisions.

Thanks



