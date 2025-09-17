const config = {
  branches: ['main'],
  plugins: [
    '@semantic-release/commit-analyzer',
    '@semantic-release/release-notes-generator',
    [
      '@semantic-release/changelog',
      {
        changelogFile: 'CHANGELOG.md',
        changelogTitle: '# Changelog',
      },
    ],
    [
      '@semantic-release/git',
      {
        assets: ['package.json', 'CHANGELOG.md'],
        message:
          'chore(release): ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}',
      },
    ],
    [
      '@semantic-release/exec',
      {
        successCmd: `
          BRANCH="release/v\${nextRelease.version}";
          git fetch origin;
          if git show-ref --verify --quiet refs/remotes/origin/$BRANCH; then
            echo "Branch $BRANCH já existe. Atualizando...";
            git checkout -B $BRANCH;
          else
            echo "Criando nova branch $BRANCH...";
            git checkout -b $BRANCH;
          fi
          git push origin $BRANCH --force;
        `,
      },
    ],
    '@semantic-release/github',
  ],
};

module.exports = config;
