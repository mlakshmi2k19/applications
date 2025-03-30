from flask import Blueprint, render_template, request
from covid_india import states

views = Blueprint('views', __name__)

@views.route('/', methods=['GET', 'POST'])
def home():
    cases = {}
    state = "Tamil Nadu"
    error = ""

    if request.method == 'POST':
        state = request.form.get('state')
        try:
            res = states.getdata(state)
            if not res or not isinstance(res, dict):
                error = "API Error: Invalid response"
            else:
                cases = {
                    'Total': res.get('Total', 'N/A'),
                    'Active': res.get('Active', 'N/A'),
                    'Cured': res.get('Cured', 'N/A'),
                    'Death': res.get('Death', 'N/A'),
                }
        except Exception as e:
            error = f"API Error: {str(e)}"

    return render_template("home.html", cases=cases, state=state, error=error)
