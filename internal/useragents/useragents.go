package useragents

// User agent component data
var (
	platforms = []string{"Macintosh", "Windows", "X11"}
	
	macOSVersions = []string{"68K", "PPC", "Intel Mac OS X"}
	
	windowsVersions = []string{
		"Win3.11", "WinNT3.51", "WinNT4.0", "Windows NT 5.0",
		"Windows NT 5.1", "Windows NT 5.2", "Windows NT 6.0",
		"Windows NT 6.1", "Windows NT 6.2", "Win 9x 4.90",
		"WindowsCE", "Windows XP", "Windows 7", "Windows 8",
		"Windows NT 10.0; Win64; x64",
	}
	
	linuxVersions = []string{"Linux i686", "Linux x86_64"}
	
	browsers = []string{"chrome", "spider", "ie"}
	
	ieOptions = []string{".NET CLR", "SV1", "Tablet PC", "Win64; IA64", "Win64; x64", "WOW64"}
	
	spiderBots = []string{
		"AdsBot-Google ( http://www.google.com/adsbot.html)",
		"Baiduspider ( http://www.baidu.com/search/spider.htm)",
		"FeedFetcher-Google; ( http://www.google.com/feedfetcher.html)",
		"Googlebot/2.1 ( http://www.googlebot.com/bot.html)",
		"Googlebot-Image/1.0",
		"Googlebot-News",
		"Googlebot-Video/1.0",
	}
)