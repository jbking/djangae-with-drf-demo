from django.http import HttpResponse
from django.conf import settings


def index(request):
    return HttpResponse("I'm a module, using %s" % settings.SETTINGS_MODULE)
