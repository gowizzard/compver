// Add labels to pull requests based on the title. This is based on the conventional commits format.
// Here you can find more information: https://www.conventionalcommits.org/en/v1.0.0/#summary.
module.exports = ({github, context}) => {

    const title = context.payload.pull_request.title.toLowerCase();

    if (!title) return;

    function getLabelFromTitle(title) {
        if (title.includes('feat!:')) return 'breaking change';
        if (title.includes('fix:')) return 'fix';
        if (title.includes('feat:')) return 'feature';
        return null;
    }

    const label = getLabelFromTitle(title);
    if (label) {
        github.rest.issues.addLabels({
            issue_number: issue.number,
            owner: repo.owner,
            repo: repo.repo,
            labels: [label]
        });
    }

};