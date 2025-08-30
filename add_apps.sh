#!/bin/bash

jobtracker add --company "Facebook" --position "Software Engineer" --status "Applied" --applied_on "2023-10-01"
jobtracker add --company "Google" --position "Data Scientist" --status "Interview" --applied_on "2023-09-15"
jobtracker add --company "Apple" --position "Machine Learning Engineer" --status "Applied" --applied_on "2023-10-01"
jobtracker add --company "Microsoft" --position "Machine Learning Specialist"
jobtracker add --company "Huawei" --position "Frontend Developer"
jobtracker add -c "Luxoft" -p "Backend Developer"
jobtracker add -c "NCR" -p "Devops"
jobtracker list
