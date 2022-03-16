var app = new Vue({
  el: '#app',
  delimiters: ['[[',']]'],
  data: {
      adding: false,
      serverError: false,
      serverMessage: "",
      message: 'Hello Vue!',
      domain: '',
      selected: 0,
      content: '',
      list: [
          {
              "language": "Curl",
              "classname": "language-shell",
              "content": `
# Try our API. Copy & run this in your terminal. 
curl -s --user 'api:d5e76748-7b9a-11e9-bb91-f21898b25098' \\\n'https://api.hedwi.com/mail/example.hedwi.com' \\\n-F from='user@example.hedwi.com' \\\n-F to='support@hedwi.com' \\\n-F subject='Hello' \\\n-F text='Testing email from Hedwi!'`,
          },
          {
              "language": "Go",
              "classname": "language-go",
              "content": `
import (
	"fmt"
	b64 "encoding/base64"
	"github.com/levigross/grequests"
)
func main() {
    auth := b64.StdEncoding.EncodeToString([]byte("api:d5e76748-7b9a-11e9-bb91-f21898b25098"))
    ro := &grequests.RequestOptions{
        Headers: map[string]string{"Authorization": "Basic " + auth},
        Data: map[string]string{ "from": "user@example.hedwi.com",   "to": "support@hedwi.com",
                                 "subject": "Hello", "text": "Testing email from Hedwi!",},
    }
    r, _ := grequests.Post("https://api.hedwi.com/mail/example.hedwi.com", ro)
    fmt.Println(r.String())
}`,
          },
          {
              "language": "Python",
              "classname": "language-python",
              "content": `
import requests

def send_email():
    return requests.post(
        "https://api.hedwi.com/mail/example.hedwi.com", 
        auth = ("api", "d5e76748-7b9a-11e9-bb91-f21898b25098"),
        data = {
            "from": "user@example.hedwi.com>",
            "to": "support@hedwi.com",
            "subject": "Hello",
            "text": "Testing email from Hedwi!"})`,
        },
        {
             "language": "PHP",
             "classname": "language-php",
             "content": `\n# Try running this locally.
function sendEmail() {
    $data = array("from" => "user@example.hedwi.com>", "to" => "support@hedwi.com",
                  "subject" => "Hello", "text" => "Testing email from Hedwi!");
    $ch = curl_init("https://api.hedwi.com/mail/example.hedwi.com");
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLINFO_HEADER_OUT, true);
    curl_setopt($ch, CURLOPT_POST, true);
    curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
    curl_setopt($ch, CURLOPT_USERPWD, "api:d5e76748-7b9a-11e9-bb91-f21898b25098");
    $result = curl_exec($ch);
    curl_close($ch);
    return $result;
}`
},
        {"language": "C#",
            "classname": "language-csharp",
            "content": `\n# Try running this locally.
public static RestResponse SendSimpleMessage() {
    RestClient client = new RestClient("https://api.hedwi.com");
    client.Authenticator = new HttpBasicAuthenticator("api","d5e76748-7b9a-11e9-bb91-f21898b25098");
    RestRequest request = new RestRequest();
    request.AddParameter("domain", "example.hedwi.com", ParameterType.UrlSegment);
    request.Resource = "mail/{domain}";
    request.AddParameter("from", "user@example.hedwi.com");
    request.AddParameter("to", "support@hedwi.com");
    request.AddParameter("subject", "Hello");
    request.AddParameter("text", "Testing email from Hedwi!");
    request.Method = Method.POST;
    return client.Execute(request);
}`},
    {
        "language": "Ruby",
        "classname": "language-ruby",
            "content": `
# Try running this locally.
def send_simple_message
  RestClient.post "https://api:d5e76748-7b9a-11e9-bb91-f21898b25098"
  "@api.hedwi.com/mail/example.hedwi.com",
  :from => "user@example.hedwi.com",
  :to => "support@hedwi.com",
  :subject => "Hello",
  :text => "Testing email from Hedwi!"
end`},
      ],
  },
    computed: {
        classObject: function() {
            return {
                is_active: this.current_tab == this.tab,
            }
        }
    },
  methods: {
        select: function(index) {
          this.selected = index;
          this.content = this.list[index]['content'];
          Prism.highlightAll();
        },
        start: function(){
            window.location.href = "/signup";
        },
        configDomain: function(item){
        },
        deleteDomain: function(item){
        },
        saveDomain: function(item){
        },
        cancleDomain: function(index) {
        },
        addDomain: function() {
        },
        getDomains: function() {
        },
    },
    created() {
    }
})
