// www.go

package www

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}
	fmt.Fprint(w, index_html, u, ih2)
}

const index_html = ` 
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
	    <title>Clio: Online Legal Practice Management Software | SaaS for Lawyers, Attorneys, Law Firms</title>
	    <meta name="description" content="Clio makes law practice management easy with legal case management software &amp; online tools for small- to mid-sized law firms." />
	    <meta name="author" content="David Skinner" />
		<meta name="robots" content="INDEX, NOFOLLOW" />
		<link rel="icon" type="image/png" href="./images/favicon.png" />
		<link rel="shortcut icon" href="./images/favicon.ico" />
		<link type="text/css" rel="stylesheet" media="all" href="./stylesheets/davsk.css" />
	</head>

	<body>
		<div id="container">
			<div style=" text-align: left; text-indent: 0px; padding: 0px 0px 0px 0px; margin: 0px 0px 0px 0px;">
				<table width="100%" border="1" cellpadding="2" cellspacing="2" style="background-color: #ffffff;">
					<tr valign="top">
						<td style="border-width : 0px;"><br />
							<img src="./images/clio_certifiedprogram.png" width="200" height="200" alt="Checkmark clio certified consultant" border="0">
						</td>

						<td style="border-width : 0px; text-align: center;"><br />
							<div id="header">
								<h1>GoClio.Davsk.com</h1>
								<p class="description">Appropriate application of advanced technology.</p>
							</div>
						</td>

						<td style="border-width : 0px; text-align: right;"><br />
							<img src="./images/davsk_logo.png" width="417" height="200" alt="" border="0">
						</td>
					</tr>
				</table>
			</div>

			<div id="wrapper">
				<div id="content">
					<h1>Hello, `

const ih2 = `!</h1>
				<h2>30 Day Free Trial</h2>
				<p>Clio leads the way in cloud-based practice management with a rich set of features that make managing everything from intake to invoicing a snap. Signup for a free trial at
				<a href="http://www.goclio.com/signup?referral_code=davsk">http://www.goclio.com/signup?referral_code=davsk</a>
				and review the features available.</p>
				<h2>Local Support</h2>
				<p>You have local support at no additional charge to assist with commisioning, importing historical data and contacts, establishing procedures, and training staff.
				Online support and toll free support are also available as part of the package.</p>
				<h2>Local Community</h2>
				<p>Local firms using GoClio are available. You can see how others have integrated their firm management. You can talk to the people who use GoClio in their practice every workday.</p>
				<h2>Product Integration</h2>
				<ul>
					<li><b>Document Management:</b> DropBox, Box, Google Drive</li>
					<li><b>Calendars and Contacts:</b>
						<ul>
							<li>Clio Sync for Outlook provides bidirectional sync between Clio and Microsoft Outlook for contacts, calendar entries and tasks.</li>
							<li>Google Apps Integration. Single Sign-On, calendar sync and contact sync all help make sure switching between Google Apps and Clio is completely painless.</li>
						</ul>
					<li><b>Email:</b> Compatible with all e-mail clients, including Microsoft Outlook, Apple Mail, Thunderbird, and Gmail</li> 
					<li><b>Accounting:</b> Quickbooks</li>
				</ul>
				<h2>Time Matters<sup>&reg;</sup></h2>
				<p>Conversions can sometimes be problematic. Data validity of prior database needs to be verified to avoid importing junk data.<br>
				I have experience with the exporting of <b>Time Matters<sup>&reg;</sup></b>
				data from <b>MsSql<sup>&reg;</sup></b>
				database for bulk import and with the <b>AutoHotKey<sup>&reg;</sup></b>
				database driven automation of <b>GoClio<sup>&reg;</sup></b> data entry to preserve historical data and with custom AI programming for data validation.</p>
				<h2>Additional Services</h2>
				<p>Services not directly related to GoClio such as certification of network wiring, configuration of routers, compliance with Gramm-Leach-Bliley Act, may be contracted at standard rates
				<a href="http://forensic.davsk.com/index_files/page0007.html">http://forensic.davsk.com/index_files/page0007.html</a>
				</p>
			</div>
		</div>

		<div id="navigation">
			<h2>Hot Links</h2>
			<ul>
				<li><a href="http://www.goclio.com/signup?referral_code=davsk">GoClio Signup</a></li>
				<li><a href="http://forensic.davsk.com">Forensic Services</a></li>
			</ul>
		</div>

		<div id="extra">
			<h2>Contact Information</h2>
			<p>David Skinner<br>
			650 S5th St Apt 88<br>
			Arkadelphia  AR   71923-6242</p>
			<p>text: <em>501-201-0414</em></p>
			<p>email: 
				<script type="text/javascript">
					/*<![CDATA[*//******* Tool from Privacy Policy Online
					* URL: http://www.PrivacyPolicyOnline.com* *************/
					<!-- Encrypted version of: your email //-->
					var CodedArray=[115,107,105,110,110,101,114,46,100,97,118,105,100,64,103,109,97,105,108,46,99,111,109,]
					var encryptedEmail='' 
					for (var i=0; i<CodedArray.length; i++)
					 encryptedEmail+=String.fromCharCode(CodedArray[i])
					document.write('<a href="mailto:')
					document.write(encryptedEmail)
					document.write('">David Skinner</a>')
					/*]]>*/
				</script>
			</p>
		</div>

			<div id="footer">
				<ul>
					<li><a href="http://goclio.davsk.com/disclaimer.html">Disclaimer</a></li>
					<li><a href="http://goclio.davsk.com/privacypolicy.html">Privacy Policy</a></li>
					<li><a href="http://goclio.davsk.com/termsofservice.html">Terms of Service</a></li>
				</ul>

				<p align=center>
					&copy; Copyright 2012 Davsk - All Rights Reserved
				</p>

				<script>
					(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
					(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
					m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
					})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
					ga('create', 'UA-52919104-2', 'auto');
					ga('send', 'pageview');
				</script>

				<a href="http://www.freecounterstat.com" title="hit counter"><img src="http://counter5.statcounterfree.com/private/freecounterstat.php?c=c1c57dddb7bcc5af4844eeb566949c75" border="0" title="hit counter" alt="hit counter"></a>
			</div>
		</div>
	</body>
</html>

`
