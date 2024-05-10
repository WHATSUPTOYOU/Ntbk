# 任务描述
实现基于selenium的网页登录配置操作，完成自动化的页面登陆和配置功能
# 实现
1. 安装 selenium
 	- `pip install selenium`
 2. 下载 web 驱动
 	- 以 Chrome 为例，根据驱动版本下载对应的浏览器驱动，下载地址可参考(https://www.cnblogs.com/aiyablog/articles/17948703)
 3. 安装 web 驱动
	 - 将驱动解压后放置到 python 环境目录下 python.exe 文件所在的目录，如C:\Users\XXX\AppData\Local\Programs\Python\Python39
4. 代码实现示例如下：
```python
import time

from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

from selenium.webdriver import DesiredCapabilities

capabilities = DesiredCapabilities.CHROME.copy()
capabilities['acceptSslCerts'] = True # 设置忽略ssl证书认证
capabilities['acceptInsecureCerts'] = True

# # 在driver中加入desired_capabilities参数
# driver = webdriver.Chrome(options=chrome_options, desired_capabilities=capabilities)

chrome_options = webdriver.ChromeOptions()
# 使用headless无界面浏览器模式
# chrome_options.add_argument('--headless') # 设置无界面模式
# chrome_options.add_argument('--disable-gpu')
chrome_options.add_argument('--ignore-certificate-errors')

if __name__ == '__main__':

    driver = webdriver.Chrome(options=chrome_options)
    driver.get("https://10.19.196.144/doc/index.html#/portal/login")
    driver.implicitly_wait(20) # 设置隐式等待，等待要搜索的元素加载完成，超时则退出
    print(f"browser text = {driver.page_source}")
    # 根据XPATH路径获取元素，直接从页面元素复制即可获得
    element = driver.find_element(By.XPATH, '//*[@id="portal"]/div/div/div[1]/div[2]/div/form/div[1]/div/div/input')
    element.send_keys("admin")
    element = driver.find_element(By.XPATH, '//*[@id="portal"]/div/div/div[1]/div[2]/div/form/div[2]/div/span/div/input')
    element.send_keys("seclab12345")
    element.send_keys(Keys.ENTER)

    element = driver.find_element(By.XPATH, '//*[@id="app"]/div[2]/div[1]/div[2]/ul/li[5]/i')
    element.click()
    element = driver.find_element(By.XPATH, '//*[@id="operations"]/div/div[1]/div/div[1]/div/div/div/div/ul/div[2]/li')
    element.click()
    element = driver.find_element(By.XPATH, '//*[@id="tab-EDRplatform"]')
    element.click()
    element = driver.find_element(By.XPATH, '//*[@id="operations"]/div/div[2]/div/div/div[2]/div/div/div/form/div[1]/div/label/span/span')
    if "2px" in element.get_attribute('style'):
        print("click")
        element.click()
    element = driver.find_element(By.XPATH, '//*[@id="operations"]/div/div[2]/div/div/div[2]/div/div/div/form/div[2]/div/div/div/input')
    element.clear()
    element.send_keys("10.19.201.78")
    element = driver.find_element(By.XPATH, '//*[@id="operations"]/div/div[2]/div/div/div[2]/div/div/div/form/div[3]/div/div/div/input')
    element.clear()
    element.send_keys("17660")
    element = driver.find_element(By.XPATH, '//*[@id="operations"]/div/div[2]/div/div/div[2]/div/div/div/form/div[4]/div/span/div/input')
    element.click()
    element.clear()
    element.send_keys("1GGMFfcD")
    element.clear()
    element.send_keys("1GGMFfcD")
    element = driver.find_element(By.XPATH, '//*[@id="operations"]/div/div[2]/div/div/div[2]/div/div/div/form/div[7]/div/div/div[1]/div/div/button')
    time.sleep(5)
    element.click()

    time.sleep(50)
    # driver.find_element(By.XPATH, "//input[@name='username']")
    driver.close()

```