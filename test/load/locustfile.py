from locust import HttpUser, task

class GoServerTest(HttpUser):
    @task
    def home(self):
        # Reach the home page
        self.client.get("/")
    
    @task
    def users(self):
        # Get the full list of users
        self.client.get("/users")

    @task
    def failures(self):
        # Get a failure
        self.client.get("/fail")

    @task
    def notFound(self):
        # Get to the 404 page
        self.client.get("/404")

    @task 
    def createUser(self):
        # Create a new user
        # current_time = str(int(round(time.time() * 1000)))
        # self.client.post("/user/create", {"Id":current_time, "First_name":current_time, "Last_name":current_time, "Email":current_time}})
        self.client.get("/user/create")
	
    @task 
    def updateUser(self):
        # current_time = str(int(round(time.time() * 1000)))
        # self.client.get("/user/update", {"Id":current_time, "First_name":current_time, "Last_name":current_time, "Email":current_time}})
        self.client.get("/user/update")
	
    @task 
    def replaceUser(self):
        # current_time = str(int(round(time.time() * 1000)))
        # self.client.get("/user/replace", {"Id":current_time, "First_name":current_time, "Last_name":current_time, "Email":current_time}})
        self.client.get("/user/replace")
	
    @task 
    def deleteUser(self):
        # current_time = str(int(round(time.time() * 1000)))
        # self.client.get("/user/delete", {"Id":current_time, "First_name":current_time, "Last_name":current_time, "Email":current_time}})
        self.client.get("/user/delete")