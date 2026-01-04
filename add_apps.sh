#!/bin/bash

jobtracker add --company "Facebook" --position "Software Engineer" --status "Applied"
jobtracker add --company "Google" --position "Data Scientist" --status "Interview"
jobtracker add --company "Apple" --position "Machine Learning Engineer" --status "Applied"
jobtracker add --company "Microsoft" --position "Machine Learning Specialist" --status "Applied"
jobtracker add --company "Huawei" --position "Frontend Developer" --status "Applied"
jobtracker add -c "Luxoft" -p "Backend Developer" -s "Applied"
jobtracker add -c "NCR" -p "Devops" -s "Applied"
jobtracker list
