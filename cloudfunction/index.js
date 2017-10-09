const got = require('got');
const url = require('url');
const mime = require('mime-types');

var gcs = require('@google-cloud/storage')({
				  projectId: 'grape-spaceship-123',
				  keyFilename: '/path/to/keyfile.json'
});

exports.github_merge = function github_merge(request, response) {
  var bucket = '';
  var branch = request.body.pull_request.base.ref;
  var merged = request.body.merged;
  var action = request.body.action;
  var pull_request_url = request.body.pull_request.url;
  var files_url = request.body.pull_request.url + '/files';
  switch(branch) {
    case 'master':
      bucket = 'darkmane-showcase.appspot.com';
      break;
    case 'staging':
      bucket = 'staging.darkmane-showcase.appspot.com';
      break;
    default:
      response.status(200).send('{}');
      return;
  }
  if (action == 'closed' && merged) {
    // get list of affected files
		var files = []
		getAffectedFiles(pull_request_url).then((files) => {
			files.forEach((file, i) => {
				// files[i] = makeRequest(file.raw_url)
				files[i] = makeRequest(file.contents_url)
					.then((contents_url) => makeRequest(contents_url))
					.then((contents) => {
						  var path = contents.path;
						  var content = contents.content;
						  var encoding = contents.encoding;
							var mimetype = mimetypes.lookup(path);
						  const file = bucket.file(path);
						  const stream = file.createWriteStream({
						  	  metadata: { contentType: mimetype }
						    }, public: true);

							stream.on('error', (err) => {

							});

							stream.on('finish', (err) => {
									req.file.cloudStorageObject = gcsname;
								  file.makePublic().then(() => {
								  req.file.cloudStoragePublicUrl = getPublicUrl(gcsname);
							});

							stream.end(contents.content);

				   	});
			}
			return Promise.all(files);
		}).then((file_promises) => {
			file_promises.forEach(fp, i) => {
				fp.then((contents) => contents)
					.then(() => {})

				}
			}
		})
    // loop through list of files
    //   get file from github
    //   push to google drive
  }
}
function makeRequest (uri, options) {
  options || (options = {});

  // Add appropriate headers
  options.headers || (options.headers = {});
  options.headers.Accept = 'application/vnd.github.black-cat-preview+json,application/vnd.github.v3+json';

  // Send and accept JSON
  options.json = true;
  if (options.body) {
    options.headers['Content-Type'] = 'application/json';
    if (typeof options.body === 'object') {
      options.body = JSON.stringify(options.body);
    }
  }

  // Add authentication
  const parts = url.parse(uri);
  parts.auth = `jmdobry:${settings.accessToken}`;

  // Make the request
  return got(parts, options).then((res) => res.body);
}

function getAffectedFiles(pr_url) {
  return makeRequest('${pr_url}/files')
}
