const express = require('express');
const router = express.Router();
const { authenticateToken } = require('../middleware/auth');
const versionController = require('../controllers/versionController');

/**
 * Version Control API Routes
 * 
 * Provides Git-like version control operations for video projects
 */

// Get commit history
router.get(
  '/projects/:projectId/commits',
  authenticateToken,
  versionController.getCommits
);

// Create new commit
router.post(
  '/projects/:projectId/commits',
  authenticateToken,
  versionController.createCommit
);

// Get specific commit
router.get(
  '/projects/:projectId/commits/:commitId',
  authenticateToken,
  versionController.getCommit
);

// List branches
router.get(
  '/projects/:projectId/branches',
  authenticateToken,
  versionController.getBranches
);

// Create branch
router.post(
  '/projects/:projectId/branches',
  authenticateToken,
  versionController.createBranch
);

// Switch branch (checkout)
router.post(
  '/projects/:projectId/checkout',
  authenticateToken,
  versionController.checkout
);

// Merge branches
router.post(
  '/projects/:projectId/merge',
  authenticateToken,
  versionController.merge
);

// Get diff between commits
router.get(
  '/projects/:projectId/diff',
  authenticateToken,
  versionController.getDiff
);

// Rollback to specific commit
router.post(
  '/projects/:projectId/rollback',
  authenticateToken,
  versionController.rollback
);

module.exports = router;
