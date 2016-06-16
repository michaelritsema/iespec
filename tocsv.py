import json
import csv

f = open('/tmp/test.json','r')
f.read(17)


first = True
w = None
for line in f.readlines():
	j = json.loads( "[" + line + "]")
	#print(j[0].keys())
	if first:
		fn = [x for x in j[0].keys()]
		fn.append('sourceIPv6Address')
		fn.append('destinationIPv6Address')
		w = csv.DictWriter(open("/tmp/pcap.csv", "wt+"), fieldnames=fn + [''])
		w.writeheader()
		first = False


	w.writerows(j)
		

