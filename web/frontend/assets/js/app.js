var application = {};

(function() {
	application.init = function() {
		application.initAdminNavbar();
		application.submitFormsOnMetaEnter();
	};

    application.initAdminNavbar = function() {
		var url = window.location.protocol + "//" + window.location.host + window.location.pathname;
		$('ul.nav a').filter(function() {
			if(this.href == url) {
                return true;
			}
			return false;
		}).parent().addClass('active');
    };

    application.submitFormsOnMetaEnter = function() {
    	document.body.addEventListener('keydown', function(e) {
            if(!(e.keyCode == 13 && (e.metaKey || e.altKey || e.ctrlKey))) return;

        	var target = e.target;
        	if(target.form) {
        		target.form.submit();
        	}
        });
    };
})()

$(document).ready(function() {
	application.init();
});