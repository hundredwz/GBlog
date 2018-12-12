package dao

var (
	drop_gblog_content_meta = `DROP TABLE IF EXISTS "gblog_content_meta"`
	gblog_content_meta      = `CREATE TABLE gblog_content_meta (
  Cid int(11) NOT NULL,
  Mid int(11) NOT NULL,
  PRIMARY KEY (Cid,Mid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	drop_gblog_contents = `DROP TABLE IF EXISTS "gblog_contents"`
	gblog_contents      = `CREATE TABLE gblog_contents (
  Cid int(11) NOT NULL AUTO_INCREMENT,
  Title varchar(200),
  Slug varchar(200) DEFAULT '',
  Created datetime,
  Modified datetime,
  Text longtext,
  Order_ int(11) DEFAULT 0,
  AuthorId int(11) DEFAULT 0,
  Template varchar(255) DEFAULT '',
  Type varchar(255) DEFAULT 'article',
  Status varchar(255) DEFAULT 'publish',
  Password varchar(255) DEFAULT '',
  CommentsNum int(11) DEFAULT 0,
  AllowComment tinyint(1) DEFAULT 1,
  AllowPing tinyint(1) DEFAULT 1,
  AllowFeed tinyint(1) DEFAULT 1,
  Parent int(11) DEFAULT 0,
  PRIMARY KEY (Cid),
  UNIQUE KEY slug (Slug)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	drop_gblog_comments = `DROP TABLE IF EXISTS "gblog_comments"`
	gblog_comments      = `CREATE TABLE gblog_comments (
  Coid int(11) NOT NULL AUTO_INCREMENT,
  Cid int(11) DEFAULT 0,
  Created datetime,
  Author varchar(255),
  AuthorId int(11) DEFAULT 0,
  OwnerId int(11) DEFAULT 0,
  Mail varchar(255) default '',
  Url varchar(255) default '',
  Ip varchar(255) default '',
  Agent varchar(255) default '',
  Text text default '',
  Type varchar(255) DEFAULT 'comment',
  Status varchar(255) DEFAULT 'approved',
  Parent int(11) DEFAULT 0,
  PRIMARY KEY (Coid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	drop_gblog_metas = `DROP TABLE IF EXISTS "gblog_metas"`
	gblog_metas      = `CREATE TABLE gblog_metas (
  Mid int(11) NOT NULL AUTO_INCREMENT,
  Name varchar(200),
  Slug varchar(200),
  Type varchar(32) NOT NULL,
  Description varchar(200),
  Count int(11) DEFAULT 0,
  Order_ int(11) DEFAULT 0,
  Parent int(11) DEFAULT 0,
  PRIMARY KEY (Mid),
  UNIQUE KEY name (Name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	drop_gblog_users = `DROP TABLE IF EXISTS "gblog_users"`
	gblog_users      = `CREATE TABLE gblog_users (
  Uid int(10) unsigned NOT NULL auto_increment,
  Name varchar(32) default '',
  Password varchar(64) default '',
  Mail varchar(200) default '',
  Url varchar(200) default '',
  ScreenName varchar(32) default '',
  Created datetime,
  Logged datetime,
  Permission varchar(64) default 'admin',
  PRIMARY KEY  (Uid),
  UNIQUE KEY (Name),
  UNIQUE KEY (Mail)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8;`
)
