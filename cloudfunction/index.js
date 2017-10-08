exports.github_merge = function github_merge(request, response) {
   var bucket = ''
	 var branch = request.body.pull_request.base.ref
	 var merged = request.body.merged
	 var action = request.body.action
	 var files_url = request.body.pull_request.url + '/files'
	 switch(branch) {
		case 'master':
		 bucket = 'darkmane-showcase.appspot.com'
		 break;
		case 'staging':
		 bucket = 'staging.darkmane-showcase.appspot.com'
		 break;
		default:
			response.status(200).send('{}');
			return;

	 }
	 var deploy = (action == 'closed' && merged)

}
