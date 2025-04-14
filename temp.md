### Create and activate a virtual environment
python3 -m venv .venv
source .venv/bin/activate  # Mac/Linux
# OR
.venv\Scripts\activate      # Windows


### installing new package
pip install seaborn

### preserve to requirements.txt so you can share it later
pip freeze > requirements.txt

### install all packages listed in requirements.txt
pip install -r requirements.txt