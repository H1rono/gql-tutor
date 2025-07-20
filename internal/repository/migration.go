package repository

// https://entgo.io/docs/versioned-migrations/

import (
	"context"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
)

type MigrateConfig struct {
	// `migrations` ディレクトリへののパス
	migrationsDir string
}

type MigrateOption func(*MigrateConfig)

// `migrations` ディレクトリへのパスを設定します.
//
// このオプションを指定しなかった場合, デフォルトでは `migrations` が使われます.
// `MigrateDiff` ないし `MigrateApply` の実行時にここで指定されたディレクトリが存在している必要があります.
func MigrationsDir(dir string) MigrateOption {
	return func(c *MigrateConfig) {
		c.migrationsDir = dir
	}
}

func defaultMigrateConfig() *MigrateConfig {
	return &MigrateConfig{"migrations"}
}

func (c *MigrateConfig) applyOptions(options ...MigrateOption) {
	for _, o := range options {
		o(c)
	}
}

func (c *MigrateConfig) load() ([]schema.MigrateOption, error) {
	dir, err := atlas.NewLocalDir(c.migrationsDir)
	if err != nil {
		return nil, err
	}
	opts := append(defaultMigrateOptions(), schema.WithDir(dir))
	return opts, nil
}

func defaultMigrateOptions() []schema.MigrateOption {
	return []schema.MigrateOption{
		schema.WithMigrationMode(schema.ModeReplay),
		schema.WithDialect(dialect.MySQL),
		schema.WithFormatter(atlas.DefaultFormatter),
	}
}

func (r *Repository) MigrateDiff(ctx context.Context, options ...MigrateOption) error {
	config := defaultMigrateConfig()
	config.applyOptions(options...)
	opts, err := config.load()
	if err != nil {
		return err
	}
	return r.c.Schema.Diff(ctx, opts...)
}

func (r *Repository) MigrateApply(ctx context.Context, options ...MigrateOption) error {
	config := defaultMigrateConfig()
	config.applyOptions(options...)
	opts, err := config.load()
	if err != nil {
		return err
	}
	return r.c.Schema.Create(ctx, opts...)
}
