//test

from selenium import webdriver
import time

chrome_options=webdriver.ChromeOptions()
chrome_options.add_argument('--no-sandbox')
driver=webdriver.Chrome('./chromedriver',chrome_options=chrome_options)

driver.set_page_load_timeout(10)
driver.get("http://test.xyz")
driver.find_element_by_link_text('Login').click()
time.sleep(3)

driver.find_element_by_name('username').send_keys("test")
driver.find_element_by_name('password').send_keys("test")
driver.find_element_by_class_name('btn-primary').click()
time.sleep(3)

driver.find_element_by_link_text('Logout').click()

time.sleep(3)
driver.quit()
