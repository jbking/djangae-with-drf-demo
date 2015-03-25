from django.conf import settings

def check_session_csrf_enabled():
    if "session_csrf.CsrfMiddleware" not in settings.MIDDLEWARE_CLASSES:
        return [ "SESSION_CSRF_DISABLED"]

    return []
check_session_csrf_enabled.messages = { "SESSION_CSRF_DISABLED" : "Please add 'session_csrf.CsrfMiddleware' to MIDDLEWARE_CLASSES" }
