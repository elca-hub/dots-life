module.exports = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    'type-enum': [2, 'always', [
      'add',
      'remove',
      'style',
      'chore',
      'fix',
      'refactor',
      'test',
      'docs',
      'update',
      'change',
      'confuse',
      'ci',
      'revert',
      'perf',
    ]],
    'subject-max-length': [2, 'always', 50],
  }
}