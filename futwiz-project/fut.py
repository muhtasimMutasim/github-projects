
from dotenv import load_dotenv; load_dotenv()
import os
import time

from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from webdriver_manager.chrome import ChromeDriverManager

# Google\ Chrome --remote-debugging-port=9222 --user-data-dir="~/ChromeProfile"

def get_driver():
    """ Function will use chrome driver """
    user_data_dir = os.environ["USER_DATA_DIR"]
    port = os.environ["CHROME_PORT"]

    base_url = "https://www.futwiz.com/en/"
    driver_location = os.environ["CHROMEDRIVER"]
    s = Service(driver_location)
    
    options = webdriver.ChromeOptions()
    options.add_argument("user-data-dir=" + user_data_dir) 
    options.add_experimental_option("debuggerAddress", f"127.0.0.1:{port}")

    # driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()), options=options)
    driver = webdriver.Chrome(service=s, options=options)
    driver.get(base_url)

    # time.sleep(4)

    # driver.close()

    print("\n\n\ndone\n\n")
    



def main():
    """ Main function. """
    
    get_driver()



if __name__ == "__main__":
    main()