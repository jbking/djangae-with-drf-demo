from django.conf.urls import patterns, url, include
from rest_framework import routers

from . import views
from . import viewsets


router = routers.DefaultRouter()
router.register(r'polls', viewsets.PollViewSet)
router.register(r'choices', viewsets.ChoiceViewSet)

urlpatterns = patterns('',
    url(r'^$', views.index),
    url(r'^api/', include(router.urls)),
    url(r'^(?P<poll_id>\d+)/$', views.detail),
    url(r'^(?P<poll_id>\d+)/results/$', views.results),
    url(r'^(?P<poll_id>\d+)/vote/$', views.vote),
)
