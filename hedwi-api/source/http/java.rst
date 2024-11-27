.. _help-http-python:

.. _http-java:


java http 发邮件
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------


.. code-block:: java


    package com.client.modules.controller;

    import com.alibaba.fastjson.JSONObject;
    import io.swagger.annotations.Api;
    import io.swagger.annotations.ApiOperation;
    import io.swagger.annotations.ApiParam;
    import org.apache.http.client.methods.HttpPost;
    import org.apache.http.entity.StringEntity;
    import org.apache.http.impl.client.BasicResponseHandler;
    import org.apache.http.impl.client.CloseableHttpClient;
    import org.apache.http.impl.client.HttpClients;
    import org.springframework.web.bind.annotation.GetMapping;
    import org.springframework.web.bind.annotation.RequestMapping;
    import org.springframework.web.bind.annotation.RestController;

    import java.util.Base64;
    import java.util.HashMap;
    import java.util.Map;

    /**
     * @author MaiDong
     */
    @RestController
    @RequestMapping("/api")
    @Api(description = "公用组件")
    public class EmailTestController {
        private static String url = "https://xxx"; // HTTP 调用 HTTP 请求地址
        private static String key = "xxx"; // HTTP 调用 HTTP API Key

        @GetMapping("/sendEmailTest")
        @ApiOperation("发送邮件测试")
        public Object sendEmailTest(@ApiParam(value = "邮箱") String mail) {
            Base64.getDecoder().decode("");
            return 1;
        }
        public static void main(String[] args) {
            String auth = Base64.getEncoder().encodeToString(("api:" + key).getBytes());
            Map<String, String> headerMap = new HashMap<>();
            headerMap.put("Authorization", "Basic " + auth);
            Map<String, String> map = new HashMap<>();
            map.put("from", "xxx"); // 企业邮箱，格式：xxx@配置的域名地址
            map.put("to", "xxx"); // 接收邮箱
            map.put("subject", "test send 123456");
            map.put("text", "Testing email from Hedwi!");
            String result = postJson(url, JSONObject.toJSONString(map), headerMap);
            System.out.println("result: "+result);
        }

        public static String postJson(String url, String jsonBody,Map<String,String> header)  {
            String result = "";
            HttpPost httpPost = new HttpPost(url);
            CloseableHttpClient httpClient = HttpClients.createDefault();
            try {
                BasicResponseHandler handler = new BasicResponseHandler();
                StringEntity entity = new StringEntity(jsonBody, "utf-8");//解决中文乱码问题
                entity.setContentEncoding("UTF-8");
                entity.setContentType("application/json");
                httpPost.setEntity(entity);
                if(header != null && header.size() > 0){
                    header.forEach((k,v)->{
                        httpPost.setHeader(k,v);
                    });
                }
                result = httpClient.execute(httpPost, handler);
                return result;
            } catch (Exception e) {
                System.out.println("请求异常："+e.getMessage());
            } finally {
                try {
                    httpClient.close();
                } catch (Exception e) {
                    System.out.println("链接关闭异常："+e.getMessage());
                }
            }
            return result;
        }
    }

        
