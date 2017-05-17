package deployment

func GetNodeFunctionFileScaffolding() string {
	return `exports.handler = function(event, context, callback) {
    // implement stuff here
    callback(null, "message_goes_here");
};
	`
}
