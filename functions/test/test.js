exports.handler = async () => {
        return {
            statusCode: 200,
            body: "aaa | " + process.env.APP_ENV
        };
};